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

type DefaultVpc struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DefaultVpcSpec   `json:"spec,omitempty"`
	Status            DefaultVpcStatus `json:"status,omitempty"`
}

type DefaultVpcSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	AssignGeneratedIpv6CIDRBlock bool `json:"assignGeneratedIpv6CIDRBlock,omitempty" tf:"assign_generated_ipv6_cidr_block,omitempty"`
	// +optional
	CidrBlock string `json:"cidrBlock,omitempty" tf:"cidr_block,omitempty"`
	// +optional
	DefaultNetworkACLID string `json:"defaultNetworkACLID,omitempty" tf:"default_network_acl_id,omitempty"`
	// +optional
	DefaultRouteTableID string `json:"defaultRouteTableID,omitempty" tf:"default_route_table_id,omitempty"`
	// +optional
	DefaultSecurityGroupID string `json:"defaultSecurityGroupID,omitempty" tf:"default_security_group_id,omitempty"`
	// +optional
	DhcpOptionsID string `json:"dhcpOptionsID,omitempty" tf:"dhcp_options_id,omitempty"`
	// +optional
	EnableClassiclink bool `json:"enableClassiclink,omitempty" tf:"enable_classiclink,omitempty"`
	// +optional
	EnableClassiclinkDNSSupport bool `json:"enableClassiclinkDNSSupport,omitempty" tf:"enable_classiclink_dns_support,omitempty"`
	// +optional
	EnableDNSHostnames bool `json:"enableDNSHostnames,omitempty" tf:"enable_dns_hostnames,omitempty"`
	// +optional
	EnableDNSSupport bool `json:"enableDNSSupport,omitempty" tf:"enable_dns_support,omitempty"`
	// +optional
	InstanceTenancy string `json:"instanceTenancy,omitempty" tf:"instance_tenancy,omitempty"`
	// +optional
	Ipv6AssociationID string `json:"ipv6AssociationID,omitempty" tf:"ipv6_association_id,omitempty"`
	// +optional
	Ipv6CIDRBlock string `json:"ipv6CIDRBlock,omitempty" tf:"ipv6_cidr_block,omitempty"`
	// +optional
	MainRouteTableID string `json:"mainRouteTableID,omitempty" tf:"main_route_table_id,omitempty"`
	// +optional
	OwnerID string `json:"ownerID,omitempty" tf:"owner_id,omitempty"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type DefaultVpcStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DefaultVpcSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DefaultVpcList is a list of DefaultVpcs
type DefaultVpcList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DefaultVpc CRD objects
	Items []DefaultVpc `json:"items,omitempty"`
}
