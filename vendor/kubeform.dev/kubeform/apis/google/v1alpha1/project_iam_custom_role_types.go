package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ProjectIamCustomRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ProjectIamCustomRoleSpec   `json:"spec,omitempty"`
	Status            ProjectIamCustomRoleStatus `json:"status,omitempty"`
}

type ProjectIamCustomRoleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// Deprecated
	Deleted bool `json:"deleted,omitempty" tf:"deleted,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	Permissions []string `json:"permissions" tf:"permissions"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	RoleID  string `json:"roleID" tf:"role_id"`
	// +optional
	Stage string `json:"stage,omitempty" tf:"stage,omitempty"`
	Title string `json:"title" tf:"title"`
}

type ProjectIamCustomRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ProjectIamCustomRoleSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ProjectIamCustomRoleList is a list of ProjectIamCustomRoles
type ProjectIamCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ProjectIamCustomRole CRD objects
	Items []ProjectIamCustomRole `json:"items,omitempty"`
}
