package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

type ComputeVPNGateway struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ComputeVPNGatewaySpec   `json:"spec,omitempty"`
	Status            ComputeVPNGatewayStatus `json:"status,omitempty"`
}

type ComputeVPNGatewaySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CreationTimestamp string `json:"creationTimestamp,omitempty" tf:"creation_timestamp,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	Name        string `json:"name" tf:"name"`
	Network     string `json:"network" tf:"network"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	SelfLink string `json:"selfLink,omitempty" tf:"self_link,omitempty"`
}

type ComputeVPNGatewayStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ComputeVPNGatewaySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ComputeVPNGatewayList is a list of ComputeVPNGateways
type ComputeVPNGatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ComputeVPNGateway CRD objects
	Items []ComputeVPNGateway `json:"items,omitempty"`
}
