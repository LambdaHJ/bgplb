// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfig) DeepCopyInto(out *BGPConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfig.
func (in *BGPConfig) DeepCopy() *BGPConfig {
	if in == nil {
		return nil
	}
	out := new(BGPConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfigList) DeepCopyInto(out *BGPConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BGPConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfigList.
func (in *BGPConfigList) DeepCopy() *BGPConfigList {
	if in == nil {
		return nil
	}
	out := new(BGPConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfigSpec) DeepCopyInto(out *BGPConfigSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfigSpec.
func (in *BGPConfigSpec) DeepCopy() *BGPConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BGPConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfigStatus) DeepCopyInto(out *BGPConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfigStatus.
func (in *BGPConfigStatus) DeepCopy() *BGPConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BGPConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfiguration) DeepCopyInto(out *BGPConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfiguration.
func (in *BGPConfiguration) DeepCopy() *BGPConfiguration {
	if in == nil {
		return nil
	}
	out := new(BGPConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfigurationList) DeepCopyInto(out *BGPConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BGPConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfigurationList.
func (in *BGPConfigurationList) DeepCopy() *BGPConfigurationList {
	if in == nil {
		return nil
	}
	out := new(BGPConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPConfigurationSpec) DeepCopyInto(out *BGPConfigurationSpec) {
	*out = *in
	if in.ServiceExternalIPs != nil {
		in, out := &in.ServiceExternalIPs, &out.ServiceExternalIPs
		*out = make([]Cidr, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPConfigurationSpec.
func (in *BGPConfigurationSpec) DeepCopy() *BGPConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(BGPConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPIPsConfig) DeepCopyInto(out *BGPIPsConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPIPsConfig.
func (in *BGPIPsConfig) DeepCopy() *BGPIPsConfig {
	if in == nil {
		return nil
	}
	out := new(BGPIPsConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPIPsConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPIPsConfigList) DeepCopyInto(out *BGPIPsConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BGPIPsConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPIPsConfigList.
func (in *BGPIPsConfigList) DeepCopy() *BGPIPsConfigList {
	if in == nil {
		return nil
	}
	out := new(BGPIPsConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BGPIPsConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPIPsConfigSpec) DeepCopyInto(out *BGPIPsConfigSpec) {
	*out = *in
	if in.IPItems != nil {
		in, out := &in.IPItems, &out.IPItems
		*out = new(IPItemList)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPIPsConfigSpec.
func (in *BGPIPsConfigSpec) DeepCopy() *BGPIPsConfigSpec {
	if in == nil {
		return nil
	}
	out := new(BGPIPsConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BGPIPsConfigStatus) DeepCopyInto(out *BGPIPsConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BGPIPsConfigStatus.
func (in *BGPIPsConfigStatus) DeepCopy() *BGPIPsConfigStatus {
	if in == nil {
		return nil
	}
	out := new(BGPIPsConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Cidr) DeepCopyInto(out *Cidr) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Cidr.
func (in *Cidr) DeepCopy() *Cidr {
	if in == nil {
		return nil
	}
	out := new(Cidr)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IPItemList) DeepCopyInto(out *IPItemList) {
	*out = *in
	if in.IPs != nil {
		in, out := &in.IPs, &out.IPs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IPItemList.
func (in *IPItemList) DeepCopy() *IPItemList {
	if in == nil {
		return nil
	}
	out := new(IPItemList)
	in.DeepCopyInto(out)
	return out
}
