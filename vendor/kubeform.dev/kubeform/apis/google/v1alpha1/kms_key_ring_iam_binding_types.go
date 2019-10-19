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

type KmsKeyRingIamBinding struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              KmsKeyRingIamBindingSpec   `json:"spec,omitempty"`
	Status            KmsKeyRingIamBindingStatus `json:"status,omitempty"`
}

type KmsKeyRingIamBindingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Etag      string   `json:"etag,omitempty" tf:"etag,omitempty"`
	KeyRingID string   `json:"keyRingID" tf:"key_ring_id"`
	Members   []string `json:"members" tf:"members"`
	Role      string   `json:"role" tf:"role"`
}

type KmsKeyRingIamBindingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *KmsKeyRingIamBindingSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// KmsKeyRingIamBindingList is a list of KmsKeyRingIamBindings
type KmsKeyRingIamBindingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of KmsKeyRingIamBinding CRD objects
	Items []KmsKeyRingIamBinding `json:"items,omitempty"`
}
