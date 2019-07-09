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

// DigitaloceanDropletSpec defines the desired state of DigitaloceanDroplet
type DigitaloceanDropletSpec struct {
	Region string `json:"region,omitempty"`
	Size   string `json:"size,omitempty"`
	Image  string `json:"image,omitempty"`
	Name   string `json:"name,omitempty"`
}

// DigitaloceanDropletStatus defines the observed state of DigitaloceanDroplet
type DigitaloceanDropletStatus struct {
	Out         string `json:"out,omitempty"`
	Initialized bool   `json:"initialized,omitempty"`
}

// +kubebuilder:object:root=true

// DigitaloceanDroplet is the Schema for the digitaloceandroplets API
type DigitaloceanDroplet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DigitaloceanDropletSpec   `json:"spec,omitempty"`
	Status DigitaloceanDropletStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DigitaloceanDropletList contains a list of DigitaloceanDroplet
type DigitaloceanDropletList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DigitaloceanDroplet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&DigitaloceanDroplet{}, &DigitaloceanDropletList{})
}
