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

type AppServiceCustomHostnameBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AppServiceCustomHostnameBindingSpec   `json:"spec,omitempty"`
	Status            AppServiceCustomHostnameBindingStatus `json:"status,omitempty"`
}

type AppServiceCustomHostnameBindingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AppServiceName    string `json:"appServiceName" tf:"app_service_name"`
	Hostname          string `json:"hostname" tf:"hostname"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
}

type AppServiceCustomHostnameBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AppServiceCustomHostnameBindingSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AppServiceCustomHostnameBindingList is a list of AppServiceCustomHostnameBindings
type AppServiceCustomHostnameBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AppServiceCustomHostnameBinding CRD objects
	Items []AppServiceCustomHostnameBinding `json:"items,omitempty"`
}
