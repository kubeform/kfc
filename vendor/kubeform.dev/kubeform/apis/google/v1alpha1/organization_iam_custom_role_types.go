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

type OrganizationIamCustomRole struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              OrganizationIamCustomRoleSpec   `json:"spec,omitempty"`
	Status            OrganizationIamCustomRoleStatus `json:"status,omitempty"`
}

type OrganizationIamCustomRoleSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// Deprecated
	Deleted bool `json:"deleted,omitempty" tf:"deleted,omitempty"`
	// +optional
	Description string `json:"description,omitempty" tf:"description,omitempty"`
	OrgID       string `json:"orgID" tf:"org_id"`
	// +kubebuilder:validation:MinItems=1
	Permissions []string `json:"permissions" tf:"permissions"`
	RoleID      string   `json:"roleID" tf:"role_id"`
	// +optional
	Stage string `json:"stage,omitempty" tf:"stage,omitempty"`
	Title string `json:"title" tf:"title"`
}

type OrganizationIamCustomRoleStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *OrganizationIamCustomRoleSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// OrganizationIamCustomRoleList is a list of OrganizationIamCustomRoles
type OrganizationIamCustomRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of OrganizationIamCustomRole CRD objects
	Items []OrganizationIamCustomRole `json:"items,omitempty"`
}
