/*
Copyright The Kubeform Authors.

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
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type PrivateDNSZone struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PrivateDNSZoneSpec   `json:"spec,omitempty"`
	Status            PrivateDNSZoneStatus `json:"status,omitempty"`
}

type PrivateDNSZoneSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	MaxNumberOfRecordSets int `json:"maxNumberOfRecordSets,omitempty" tf:"max_number_of_record_sets,omitempty"`
	// +optional
	MaxNumberOfVirtualNetworkLinks int `json:"maxNumberOfVirtualNetworkLinks,omitempty" tf:"max_number_of_virtual_network_links,omitempty"`
	// +optional
	MaxNumberOfVirtualNetworkLinksWithRegistration int    `json:"maxNumberOfVirtualNetworkLinksWithRegistration,omitempty" tf:"max_number_of_virtual_network_links_with_registration,omitempty"`
	Name                                           string `json:"name" tf:"name"`
	// +optional
	NumberOfRecordSets int    `json:"numberOfRecordSets,omitempty" tf:"number_of_record_sets,omitempty"`
	ResourceGroupName  string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type PrivateDNSZoneStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *PrivateDNSZoneSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PrivateDNSZoneList is a list of PrivateDNSZones
type PrivateDNSZoneList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PrivateDNSZone CRD objects
	Items []PrivateDNSZone `json:"items,omitempty"`
}
