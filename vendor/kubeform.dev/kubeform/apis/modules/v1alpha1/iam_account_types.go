package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type IamAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IamAccountSpec   `json:"spec,omitempty"`
	Status            IamAccountStatus `json:"status,omitempty"`
}

type IamAccountSpec struct {
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	Source      string                     `json:"source" tf:"source"`

	// +optional
	// AWS IAM account alias for this account
	AccountAlias string `json:"accountAlias,omitempty" tf:"account_alias,omitempty"`
	// +optional
	// Whether to allow users to change their own password
	AllowUsersToChangePassword bool `json:"allowUsersToChangePassword,omitempty" tf:"allow_users_to_change_password,omitempty"`
	// +optional
	// Whether to create AWS IAM account password policy
	CreateAccountPasswordPolicy bool `json:"createAccountPasswordPolicy,omitempty" tf:"create_account_password_policy,omitempty"`
	// +optional
	// Whether to get AWS account ID, User ID, and ARN in which Terraform is authorized
	GetCallerIdentity bool `json:"getCallerIdentity,omitempty" tf:"get_caller_identity,omitempty"`
	// +optional
	// Whether users are prevented from setting a new password after their password has expired (i.e. require administrator reset)
	HardExpiry bool `json:"hardExpiry,omitempty" tf:"hard_expiry,omitempty"`
	// +optional
	// The number of days that an user password is valid.
	MaxPasswordAge json.Number `json:"maxPasswordAge,omitempty" tf:"max_password_age,omitempty"`
	// +optional
	// Minimum length to require for user passwords
	MinimumPasswordLength json.Number `json:"minimumPasswordLength,omitempty" tf:"minimum_password_length,omitempty"`
	// +optional
	// The number of previous passwords that users are prevented from reusing
	PasswordReusePrevention json.Number `json:"passwordReusePrevention,omitempty" tf:"password_reuse_prevention,omitempty"`
	// +optional
	// Whether to require lowercase characters for user passwords
	RequireLowercaseCharacters bool `json:"requireLowercaseCharacters,omitempty" tf:"require_lowercase_characters,omitempty"`
	// +optional
	// Whether to require numbers for user passwords
	RequireNumbers bool `json:"requireNumbers,omitempty" tf:"require_numbers,omitempty"`
	// +optional
	// Whether to require symbols for user passwords
	RequireSymbols bool `json:"requireSymbols,omitempty" tf:"require_symbols,omitempty"`
	// +optional
	// Whether to require uppercase characters for user passwords
	RequireUppercaseCharacters bool `json:"requireUppercaseCharacters,omitempty" tf:"require_uppercase_characters,omitempty"`
}

type IamAccountOutput struct {
	// The AWS Account ID number of the account that owns or contains the calling entity
	ThisCallerIdentityAccountID string `json:"thisCallerIdentityAccountID" tf:"this_caller_identity_account_id"`
	// The AWS ARN associated with the calling entity
	ThisCallerIdentityArn string `json:"thisCallerIdentityArn" tf:"this_caller_identity_arn"`
	// The unique identifier of the calling entity
	ThisCallerIdentityUserID string `json:"thisCallerIdentityUserID" tf:"this_caller_identity_user_id"`
	// Indicates whether passwords in the account expire. Returns true if max_password_age contains a value greater than 0. Returns false if it is 0 or not present.
	ThisIamAccountPasswordPolicyExpirePasswords string `json:"thisIamAccountPasswordPolicyExpirePasswords" tf:"this_iam_account_password_policy_expire_passwords"`
}

type IamAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *IamAccountOutput `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IamAccountList is a list of IamAccounts
type IamAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of IamAccount CRD objects
	Items []IamAccount `json:"items,omitempty"`
}
