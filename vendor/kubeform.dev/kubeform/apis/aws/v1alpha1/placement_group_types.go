package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type PlacementGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              PlacementGroupSpec   `json:"spec,omitempty"`
	Status            PlacementGroupStatus `json:"status,omitempty"`
}

type PlacementGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Name     string `json:"name" tf:"name"`
	Strategy string `json:"strategy" tf:"strategy"`
}

type PlacementGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// PlacementGroupList is a list of PlacementGroups
type PlacementGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of PlacementGroup CRD objects
	Items []PlacementGroup `json:"items,omitempty"`
}
