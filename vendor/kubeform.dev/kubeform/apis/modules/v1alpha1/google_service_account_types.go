package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type GoogleServiceAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleServiceAccountSpec   `json:"spec,omitempty"`
	Status            GoogleServiceAccountStatus `json:"status,omitempty"`
}

type GoogleServiceAccountSpec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// If assigning billing role, specificy a billing account (default is to assign at the organizational level).
	BillingAccountID string `json:"billingAccountID,omitempty" tf:"billing_account_id,omitempty"`
	// +optional
	// Generate keys for service accounts.
	GenerateKeys bool `json:"generateKeys,omitempty" tf:"generate_keys,omitempty"`
	// +optional
	// Grant billing user role.
	GrantBillingRole bool `json:"grantBillingRole,omitempty" tf:"grant_billing_role,omitempty"`
	// +optional
	// Grant roles for shared VPC management.
	GrantXpnRoles bool `json:"grantXpnRoles,omitempty" tf:"grant_xpn_roles,omitempty"`
	// +optional
	// Names of the service accounts to create.
	Names []string `json:"names,omitempty" tf:"names,omitempty"`
	// +optional
	// Id of the organization for org-level roles.
	OrgID string `json:"orgID,omitempty" tf:"org_id,omitempty"`
	// +optional
	// Prefix applied to service account names.
	Prefix string `json:"prefix,omitempty" tf:"prefix,omitempty"`
	// +optional
	// Project id where service account will be created.
	ProjectID string `json:"projectID,omitempty" tf:"project_id,omitempty"`
	// +optional
	// Common roles to apply to all service accounts, project=>role as elements.
	ProjectRoles []string `json:"projectRoles,omitempty" tf:"project_roles,omitempty"`
}

type GoogleServiceAccountOutput struct {
	// Service account email (for single use).
	Email string `json:"email" tf:"email"`
	// Service account emails.
	Emails string `json:"emails" tf:"emails"`
	// Service account emails.
	EmailsList string `json:"emailsList" tf:"emails_list"`
	// IAM-format service account email (for single use).
	IamEmail string `json:"iamEmail" tf:"iam_email"`
	// IAM-format service account emails.
	IamEmails string `json:"iamEmails" tf:"iam_emails"`
	// IAM-format service account emails.
	IamEmailsList string `json:"iamEmailsList" tf:"iam_emails_list"`
	// Service account key (for single use).
	Key string `json:"key" tf:"key"`
	// Map of service account keys.
	Keys string `json:"keys" tf:"keys"`
	// Service account resource (for single use).
	ServiceAccount string `json:"serviceAccount" tf:"service_account"`
	// Service account resources.
	ServiceAccounts string `json:"serviceAccounts" tf:"service_accounts"`
}

type GoogleServiceAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GoogleServiceAccountOutput `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleServiceAccountList is a list of GoogleServiceAccounts
type GoogleServiceAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleServiceAccount CRD objects
	Items []GoogleServiceAccount `json:"items,omitempty"`
}