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

type ServiceAccountIamMember struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ServiceAccountIamMemberSpec   `json:"spec,omitempty"`
	Status            ServiceAccountIamMemberStatus `json:"status,omitempty"`
}

type ServiceAccountIamMemberSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Etag             string `json:"etag,omitempty" tf:"etag,omitempty"`
	Member           string `json:"member" tf:"member"`
	Role             string `json:"role" tf:"role"`
	ServiceAccountID string `json:"serviceAccountID" tf:"service_account_id"`
}

type ServiceAccountIamMemberStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ServiceAccountIamMemberSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ServiceAccountIamMemberList is a list of ServiceAccountIamMembers
type ServiceAccountIamMemberList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ServiceAccountIamMember CRD objects
	Items []ServiceAccountIamMember `json:"items,omitempty"`
}
