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

type LbNATPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              LbNATPoolSpec   `json:"spec,omitempty"`
	Status            LbNATPoolStatus `json:"status,omitempty"`
}

type LbNATPoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	BackendPort int `json:"backendPort" tf:"backend_port"`
	// +optional
	FrontendIPConfigurationID   string `json:"frontendIPConfigurationID,omitempty" tf:"frontend_ip_configuration_id,omitempty"`
	FrontendIPConfigurationName string `json:"frontendIPConfigurationName" tf:"frontend_ip_configuration_name"`
	FrontendPortEnd             int    `json:"frontendPortEnd" tf:"frontend_port_end"`
	FrontendPortStart           int    `json:"frontendPortStart" tf:"frontend_port_start"`
	LoadbalancerID              string `json:"loadbalancerID" tf:"loadbalancer_id"`
	// +optional
	// Deprecated
	Location          string `json:"location,omitempty" tf:"location,omitempty"`
	Name              string `json:"name" tf:"name"`
	Protocol          string `json:"protocol" tf:"protocol"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type LbNATPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *LbNATPoolSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// LbNATPoolList is a list of LbNATPools
type LbNATPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of LbNATPool CRD objects
	Items []LbNATPool `json:"items,omitempty"`
}
