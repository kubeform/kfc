/*
Copyright The Kubeform Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by Kubeform. DO NOT EDIT.

package v1alpha1

import (
	base "kubeform.dev/kubeform/apis/base/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type FrontdoorFirewallPolicy struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FrontdoorFirewallPolicySpec   `json:"spec,omitempty"`
	Status            FrontdoorFirewallPolicyStatus `json:"status,omitempty"`
}

type FrontdoorFirewallPolicySpecCustomRuleMatchCondition struct {
	// +kubebuilder:validation:MaxItems=100
	MatchValues   []string `json:"matchValues" tf:"match_values"`
	MatchVariable string   `json:"matchVariable" tf:"match_variable"`
	// +optional
	NegationCondition bool   `json:"negationCondition,omitempty" tf:"negation_condition,omitempty"`
	Operator          string `json:"operator" tf:"operator"`
	// +optional
	Selector string `json:"selector,omitempty" tf:"selector,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=5
	Transforms []string `json:"transforms,omitempty" tf:"transforms,omitempty"`
}

type FrontdoorFirewallPolicySpecCustomRule struct {
	Action string `json:"action" tf:"action"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	MatchCondition []FrontdoorFirewallPolicySpecCustomRuleMatchCondition `json:"matchCondition,omitempty" tf:"match_condition,omitempty"`
	Name           string                                                `json:"name" tf:"name"`
	// +optional
	Priority int64 `json:"priority,omitempty" tf:"priority,omitempty"`
	// +optional
	RateLimitDurationInMinutes int64 `json:"rateLimitDurationInMinutes,omitempty" tf:"rate_limit_duration_in_minutes,omitempty"`
	// +optional
	RateLimitThreshold int64  `json:"rateLimitThreshold,omitempty" tf:"rate_limit_threshold,omitempty"`
	Type               string `json:"type" tf:"type"`
}

type FrontdoorFirewallPolicySpecManagedRuleExclusion struct {
	MatchVariable string `json:"matchVariable" tf:"match_variable"`
	Operator      string `json:"operator" tf:"operator"`
	Selector      string `json:"selector" tf:"selector"`
}

type FrontdoorFirewallPolicySpecManagedRuleOverrideExclusion struct {
	MatchVariable string `json:"matchVariable" tf:"match_variable"`
	Operator      string `json:"operator" tf:"operator"`
	Selector      string `json:"selector" tf:"selector"`
}

type FrontdoorFirewallPolicySpecManagedRuleOverrideRuleExclusion struct {
	MatchVariable string `json:"matchVariable" tf:"match_variable"`
	Operator      string `json:"operator" tf:"operator"`
	Selector      string `json:"selector" tf:"selector"`
}

type FrontdoorFirewallPolicySpecManagedRuleOverrideRule struct {
	Action string `json:"action" tf:"action"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Exclusion []FrontdoorFirewallPolicySpecManagedRuleOverrideRuleExclusion `json:"exclusion,omitempty" tf:"exclusion,omitempty"`
	RuleID    string                                                        `json:"ruleID" tf:"rule_id"`
}

type FrontdoorFirewallPolicySpecManagedRuleOverride struct {
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Exclusion []FrontdoorFirewallPolicySpecManagedRuleOverrideExclusion `json:"exclusion,omitempty" tf:"exclusion,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1000
	Rule          []FrontdoorFirewallPolicySpecManagedRuleOverrideRule `json:"rule,omitempty" tf:"rule,omitempty"`
	RuleGroupName string                                               `json:"ruleGroupName" tf:"rule_group_name"`
}

type FrontdoorFirewallPolicySpecManagedRule struct {
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Exclusion []FrontdoorFirewallPolicySpecManagedRuleExclusion `json:"exclusion,omitempty" tf:"exclusion,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	Override []FrontdoorFirewallPolicySpecManagedRuleOverride `json:"override,omitempty" tf:"override,omitempty"`
	Type     string                                           `json:"type" tf:"type"`
	Version  string                                           `json:"version" tf:"version"`
}

type FrontdoorFirewallPolicySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	// +optional
	CustomBlockResponseBody string `json:"customBlockResponseBody,omitempty" tf:"custom_block_response_body,omitempty"`
	// +optional
	CustomBlockResponseStatusCode int64 `json:"customBlockResponseStatusCode,omitempty" tf:"custom_block_response_status_code,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	CustomRule []FrontdoorFirewallPolicySpecCustomRule `json:"customRule,omitempty" tf:"custom_rule,omitempty"`
	// +optional
	Enabled bool `json:"enabled,omitempty" tf:"enabled,omitempty"`
	// +optional
	FrontendEndpointIDS []string `json:"frontendEndpointIDS,omitempty" tf:"frontend_endpoint_ids,omitempty"`
	// +optional
	Location string `json:"location,omitempty" tf:"location,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=100
	ManagedRule []FrontdoorFirewallPolicySpecManagedRule `json:"managedRule,omitempty" tf:"managed_rule,omitempty"`
	// +optional
	Mode string `json:"mode,omitempty" tf:"mode,omitempty"`
	Name string `json:"name" tf:"name"`
	// +optional
	RedirectURL       string `json:"redirectURL,omitempty" tf:"redirect_url,omitempty"`
	ResourceGroupName string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type FrontdoorFirewallPolicyStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *FrontdoorFirewallPolicySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// FrontdoorFirewallPolicyList is a list of FrontdoorFirewallPolicys
type FrontdoorFirewallPolicyList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of FrontdoorFirewallPolicy CRD objects
	Items []FrontdoorFirewallPolicy `json:"items,omitempty"`
}
