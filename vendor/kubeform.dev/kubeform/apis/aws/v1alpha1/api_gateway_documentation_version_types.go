package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiGatewayDocumentationVersion struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayDocumentationVersionSpec   `json:"spec,omitempty"`
	Status            ApiGatewayDocumentationVersionStatus `json:"status,omitempty"`
}

type ApiGatewayDocumentationVersionSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	RestAPIID   string `json:"restAPIID" tf:"rest_api_id"`
	Version     string `json:"version" tf:"version"`
}

type ApiGatewayDocumentationVersionStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayDocumentationVersionList is a list of ApiGatewayDocumentationVersions
type ApiGatewayDocumentationVersionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayDocumentationVersion CRD objects
	Items []ApiGatewayDocumentationVersion `json:"items,omitempty"`
}
