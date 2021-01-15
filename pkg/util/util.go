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

package util

import (
	corev1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}

func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}

// IsDeletionCandidate checks if object is candidate to be deleted
func IsDeletionCandidate(obj v1.Object, finalizer string) bool {
	return obj.GetDeletionTimestamp() != nil && ContainsString(obj.GetFinalizers(), finalizer)
}

// NeedToAddFinalizer checks if need to add finalizer to object
func NeedToAddFinalizer(obj v1.Object, finalizer string) bool {
	return obj.GetDeletionTimestamp() == nil && !ContainsString(obj.GetFinalizers(), finalizer)
}

// IsNeedReleaseIP checks if need to release ip.
func IsNeedReleaseIP(obj *corev1.Service, del bool) bool {
	if len(obj.Status.LoadBalancer.Ingress) > 0 {
		if obj.Spec.LoadBalancerIP != "" &&
			obj.Spec.LoadBalancerIP != obj.Status.LoadBalancer.Ingress[0].IP {
			return true
		}
		if del || obj.Spec.Type != corev1.ServiceTypeLoadBalancer {
			return true
		}
	}

	return false
}

// IsNeedAssignIP checks if need assign a new ip to object.
func IsNeedAssignIP(obj *corev1.Service) bool {
	if len(obj.Status.LoadBalancer.Ingress) == 0 && obj.Spec.Type == corev1.ServiceTypeLoadBalancer {
		return true
	}
	return false
}
