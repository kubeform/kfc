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

type DefaultNetworkACL struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultNetworkACLSpec   `json:"spec,omitempty"`
	Status            DefaultNetworkACLStatus `json:"status,omitempty"`
}

type DefaultNetworkACLSpecEgress struct {
	Action string `json:"action" tf:"action"`
	// +optional
	CidrBlock string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`
	FromPort  int    `json:"fromPort" tf:"from_port"`
	// +optional
	IcmpCode int `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`
	// +optional
	IcmpType int `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`
	// +optional
	Ipv6CIDRBlock string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`
	Protocol      string `json:"protocol" tf:"protocol"`
	RuleNo        int    `json:"ruleNo" tf:"rule_no"`
	ToPort        int    `json:"toPort" tf:"to_port"`
}

type DefaultNetworkACLSpecIngress struct {
	Action string `json:"action" tf:"action"`
	// +optional
	CidrBlock string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`
	FromPort  int    `json:"fromPort" tf:"from_port"`
	// +optional
	IcmpCode int `json:"icmpCode,omitempty" tf:"icmp_code,omitempty"`
	// +optional
	IcmpType int `json:"icmpType,omitempty" tf:"icmp_type,omitempty"`
	// +optional
	Ipv6CIDRBlock string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`
	Protocol      string `json:"protocol" tf:"protocol"`
	RuleNo        int    `json:"ruleNo" tf:"rule_no"`
	ToPort        int    `json:"toPort" tf:"to_port"`
}

type DefaultNetworkACLSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DefaultNetworkACLID string `json:"defaultNetworkACLID" tf:"default_network_acl_id"`
	// +optional
	Egress []DefaultNetworkACLSpecEgress `json:"egress,omitempty" tf:"egress,omitempty"`
	// +optional
	Ingress []DefaultNetworkACLSpecIngress `json:"ingress,omitempty" tf:"ingress,omitempty"`
	// +optional
	OwnerID string `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	SubnetIDS []string `json:"subnetIDS,omitempty" tf:"subnet_ids,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	VpcID string `json:"vpcID,omitempty" tf:"vpc_id,omitempty"`
}

type DefaultNetworkACLStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DefaultNetworkACLSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultNetworkACLList is a list of DefaultNetworkACLs
type DefaultNetworkACLList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultNetworkACL CRD objects
	Items []DefaultNetworkACL `json:"items,omitempty"`
}
