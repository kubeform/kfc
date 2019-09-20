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

type IAmAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              IAmAccountSpec   `json:"spec,omitempty"`
	Status            IAmAccountStatus `json:"status,omitempty"`
}

type IAmAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`
	Source      string                    `json:"source" tf:"source"`
	// Whether to require numbers for user passwords
	RequireNumbers bool `json:"requireNumbers" tf:"require_numbers,omitempty"`
	// Whether to require symbols for user passwords
	RequireSymbols bool `json:"requireSymbols" tf:"require_symbols,omitempty"`
	// AWS IAM account alias for this account
	AccountAlias string `json:"accountAlias" tf:"account_alias,omitempty"`
	// Whether to create AWS IAM account password policy
	CreateAccountPasswordPolicy bool `json:"createAccountPasswordPolicy" tf:"create_account_password_policy,omitempty"`
	// The number of days that an user password is valid.
	MaxPasswordAge json.Number `json:"maxPasswordAge" tf:"max_password_age,omitempty"`
	// Minimum length to require for user passwords
	MinimumPasswordLength json.Number `json:"minimumPasswordLength" tf:"minimum_password_length,omitempty"`
	// Whether to allow users to change their own password
	AllowUsersToChangePassword bool `json:"allowUsersToChangePassword" tf:"allow_users_to_change_password,omitempty"`
	// The number of previous passwords that users are prevented from reusing
	PasswordReusePrevention json.Number `json:"passwordReusePrevention" tf:"password_reuse_prevention,omitempty"`
	// Whether to get AWS account ID, User ID, and ARN in which Terraform is authorized
	GetCallerIdentity bool `json:"getCallerIdentity" tf:"get_caller_identity,omitempty"`
	// Whether users are prevented from setting a new password after their password has expired (i.e. require administrator reset)
	HardExpiry bool `json:"hardExpiry" tf:"hard_expiry,omitempty"`
	// Whether to require lowercase characters for user passwords
	RequireLowercaseCharacters bool `json:"requireLowercaseCharacters" tf:"require_lowercase_characters,omitempty"`
	// Whether to require uppercase characters for user passwords
	RequireUppercaseCharacters bool `json:"requireUppercaseCharacters" tf:"require_uppercase_characters,omitempty"`
}

type IAmAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	State  string           `json:"state,omitempty"`
	Output IAmAccountOutput `json:"output,omitempty"`
}

type IAmAccountOutput struct {
	// The AWS Account ID number of the account that owns or contains the calling entity
	ThisCallerIdentityAccountID string `json:"thisCallerIdentityAccountID" tf:"this_caller_identity_account_id"`
	// The AWS ARN associated with the calling entity
	ThisCallerIdentityArn string `json:"thisCallerIdentityArn" tf:"this_caller_identity_arn"`
	// The unique identifier of the calling entity
	ThisCallerIdentityUserID string `json:"thisCallerIdentityUserID" tf:"this_caller_identity_user_id"`
	// Indicates whether passwords in the account expire. Returns true if max_password_age contains a value greater than 0. Returns false if it is 0 or not present.
	ThisIamAccountPasswordPolicyExpirePasswords string `json:"thisIamAccountPasswordPolicyExpirePasswords" tf:"this_iam_account_password_policy_expire_passwords"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// IAmAccountList is a list of IAmAccount
type IAmAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Volume CRD objects
	Items []IAmAccount `json:"items,omitempty"`
}
