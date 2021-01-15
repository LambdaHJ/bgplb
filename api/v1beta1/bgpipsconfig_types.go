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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// BGPIPsConfigSpec defines the desired state of BGPIPsConfig
type BGPIPsConfigSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Cidr is IpRange. Edit BGPIPsConfig_types.go to remove/update
	Cidr    string      `json:"cidr"`
	Free    uint        `json:"free,omitempty"`
	Used    uint        `json:"used,omitempty"`
	IPItems *IPItemList `json:"ipItemList,omitempty"`
}

type IPItemList struct {
	IPs []string `json:"ips,omitempty"`
}

func (ipl *IPItemList) IsInUsed(ip string) bool {

	return false
}

// BGPIPsConfigStatus defines the observed state of BGPIPsConfig
type BGPIPsConfigStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// BGPIPsConfig is the Schema for the bgpipsconfigs API
type BGPIPsConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   BGPIPsConfigSpec   `json:"spec,omitempty"`
	Status BGPIPsConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// BGPIPsConfigList contains a list of BGPIPsConfig
type BGPIPsConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []BGPIPsConfig `json:"items"`
}

func init() {
	SchemeBuilder.Register(&BGPIPsConfig{}, &BGPIPsConfigList{})
}
