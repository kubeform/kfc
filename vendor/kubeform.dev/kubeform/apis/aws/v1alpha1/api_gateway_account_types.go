package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"kubeform.dev/kubeform/apis"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type ApiGatewayAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayAccountSpec   `json:"spec,omitempty"`
	Status            ApiGatewayAccountStatus `json:"status,omitempty"`
}

type ApiGatewayAccountSpecThrottleSettings struct {
	// +optional
	BurstLimit int `json:"burstLimit,omitempty" tf:"burst_limit,omitempty"`
	// +optional
	RateLimit json.Number `json:"rateLimit,omitempty" tf:"rate_limit,omitempty"`
}

type ApiGatewayAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CloudwatchRoleArn string `json:"cloudwatchRoleArn,omitempty" tf:"cloudwatch_role_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ThrottleSettings []ApiGatewayAccountSpecThrottleSettings `json:"throttleSettings,omitempty" tf:"throttle_settings,omitempty"`
}

type ApiGatewayAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayAccountSpec `json:"output,omitempty"`
	// +optional
	State *apis.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayAccountList is a list of ApiGatewayAccounts
type ApiGatewayAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayAccount CRD objects
	Items []ApiGatewayAccount `json:"items,omitempty"`
}
