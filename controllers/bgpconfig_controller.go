/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"sync"

	"github.com/LambdaHJ/bgplb/api/v1beta1"
	"github.com/LambdaHJ/bgplb/pkg/ipam"
	"github.com/LambdaHJ/bgplb/pkg/util"
	"github.com/LambdaHJ/bgplb/pkg/validate"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/controller/controllerutil"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
)

const finalizer string = "finalizer.lb.lambdahj.site/v1beta1"

const listPageSize = 50

// BGPConfigReconciler reconciles a BGPConfig object
type BGPConfigReconciler struct {
	Locker sync.Mutex
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
	ipam   *ipam.IPAMManager
}

func (r *BGPConfigReconciler) Init(reader client.Reader) error {
	ctx := context.Background()
	reqLog := r.Log.WithValues("init", "BGPConfigReconciler")

	r.ipam = ipam.NewIPAMManager()
	bgpConf := &v1beta1.BGPConfiguration{}
	nq := types.NamespacedName{Name: "default"}
	err := reader.Get(ctx, nq, bgpConf)
	if err != nil {
		if errors.IsNotFound(err) {
			return nil
		}
		return err
	}

	for _, cidr := range bgpConf.Spec.ServiceExternalIPs {
		if err := r.ipam.NewCidr(cidr.Cidr); err != nil {
			reqLog.Error(err, "creat cidr error")
		}
	}

	svcs := &corev1.ServiceList{}
	filterOptions := &client.ListOptions{Limit: listPageSize}
	for {
		err = reader.List(ctx, svcs, filterOptions)
		if err != nil {
			reqLog.Error(err, "List service error")
			return err
		}
		for _, item := range svcs.Items {
			if item.Spec.Type != corev1.ServiceTypeLoadBalancer {
				continue
			}

			r.ipam.AddUsedIP(item.Status.LoadBalancer.Ingress[0].IP)
			reqLog.Info("add used ip", "ip", item.Status.LoadBalancer.Ingress[0].IP, "service", item.Name)
		}

		if svcs.Continue == "" {
			break
		}
		filterOptions.Continue = svcs.Continue
		svcs.Continue = ""
	}
	reqLog.Info("Contriller init success")

	return nil
}

// +kubebuilder:rbac:groups=lb.lambdahj.site,resources=services,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=lb.lambdahj.site,resources=services/status,verbs=get;update;patch

func (r *BGPConfigReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	r.Locker.Lock()
	defer r.Locker.Unlock()
	ctx := context.Background()
	reqLog := r.Log.WithValues("bgpconfig", req.NamespacedName)

	svc := &corev1.Service{}
	err := r.Get(ctx, req.NamespacedName, svc)
	if err != nil {
		if errors.IsNotFound(err) {
			return ctrl.Result{}, nil
		}
		return ctrl.Result{}, err
	}

	if util.IsDeletionCandidate(svc, finalizer) {
		if util.IsNeedReleaseIP(svc, true) {
			r.ipam.ReleaseIP(svc.Status.LoadBalancer.Ingress[0].IP)
			reqLog.Info("remove ip", "ip", svc.Status.LoadBalancer.Ingress[0].IP)
		}
		controllerutil.RemoveFinalizer(svc, finalizer)
		svc.Status.LoadBalancer.Ingress = nil
		err = r.Update(ctx, svc)
		reqLog.Info("RemoveFinalizer", "finalizer", svc.Finalizers, "err", err)
		return ctrl.Result{}, err
	}

	if util.NeedToAddFinalizer(svc, finalizer) {
		controllerutil.AddFinalizer(svc, finalizer)
		err := r.Update(context.Background(), svc)
		reqLog.Info("AddFinalizer", "finalizer", svc.Finalizers, "err", err)
		if err != nil {
			return ctrl.Result{}, err
		}
	}

	if util.IsNeedReleaseIP(svc, false) {
		r.ipam.ReleaseIP(svc.Status.LoadBalancer.Ingress[0].IP)
		reqLog.Info("remove ip", "ip", svc.Status.LoadBalancer.Ingress[0].IP)
		svc.Status.LoadBalancer.Ingress = nil
	}

	if !util.IsNeedAssignIP(svc) {
		return ctrl.Result{}, nil
	}

	var ip string
	if svc.Spec.LoadBalancerIP != "" {
		reqLog.Info("specific ip", "ip", svc.Spec.LoadBalancerIP)
		if !r.ipam.AcquireSpecificIP(svc.Spec.LoadBalancerIP) {
			reqLog.Info("get specific ip error")
			return ctrl.Result{}, nil
		}
		ip = svc.Spec.LoadBalancerIP

	} else {
		ip, err = r.ipam.AcquireIP()
		if err != nil {
			reqLog.Error(err, "acquire ip error")
			return ctrl.Result{}, nil
		}
	}

	svc.Status.LoadBalancer.Ingress = append(svc.Status.LoadBalancer.Ingress, corev1.LoadBalancerIngress{IP: ip})

	err = r.Status().Update(context.Background(), svc)
	reqLog.Info("Assign exterinal IP", "IP", ip)
	if err != nil {
		return ctrl.Result{}, err
	}

	return ctrl.Result{}, nil
}

func (r *BGPConfigReconciler) SetupWithManager(mgr ctrl.Manager) error {
	p := predicate.Funcs{
		DeleteFunc: func(e event.DeleteEvent) bool {
			return false
		},
		UpdateFunc: func(e event.UpdateEvent) bool {
			return validate.IsTypeLoadBalancer(e.ObjectNew) || validate.IsTypeLoadBalancer(e.ObjectOld)
		},
		CreateFunc: func(e event.CreateEvent) bool {
			return validate.IsTypeLoadBalancer(e.Object) && !validate.IsAssignend(e.Object)
		},
	}
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.Service{}).WithEventFilter(p).
		Complete(r)
}
