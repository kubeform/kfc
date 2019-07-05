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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// LinodeInstanceSpec defines the desired state of LinodeInstance
type LinodeInstanceSpec struct {
	Region string `json:"region,omitempty"`
	Type   string `json:"type,omitempty"`
	Image  string `json:"image,omitempty"`
	Label  string `json:"label,omitempty"`
}

// LinodeInstanceStatus defines the observed state of LinodeInstance
type LinodeInstanceStatus struct {
	Out         string `json:"out,omitempty"`
	Initialized bool   `json:"initialized,omitempty"`
}

// +kubebuilder:object:root=true

// LinodeInstance is the Schema for the linodeinstances API
type LinodeInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LinodeInstanceSpec   `json:"spec,omitempty"`
	Status LinodeInstanceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LinodeInstanceList contains a list of LinodeInstance
type LinodeInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LinodeInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LinodeInstance{}, &LinodeInstanceList{})
}
