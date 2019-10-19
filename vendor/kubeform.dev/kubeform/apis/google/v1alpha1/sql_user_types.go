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

type SqlUser struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              SqlUserSpec   `json:"spec,omitempty"`
	Status            SqlUserStatus `json:"status,omitempty"`
}

type SqlUserSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	// +optional
	Host     string `json:"host,omitempty" tf:"host,omitempty"`
	Instance string `json:"instance" tf:"instance"`
	Name     string `json:"name" tf:"name"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
}

type SqlUserStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *SqlUserSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// SqlUserList is a list of SqlUsers
type SqlUserList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of SqlUser CRD objects
	Items []SqlUser `json:"items,omitempty"`
}
