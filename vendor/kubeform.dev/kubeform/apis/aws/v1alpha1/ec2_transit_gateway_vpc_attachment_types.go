package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type Ec2TransitGatewayVpcAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              Ec2TransitGatewayVpcAttachmentSpec   `json:"spec,omitempty"`
	Status            Ec2TransitGatewayVpcAttachmentStatus `json:"status,omitempty"`
}

type Ec2TransitGatewayVpcAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	DnsSupport string `json:"dnsSupport,omitempty" tf:"dns_support,omitempty"`
	// +optional
	Ipv6Support string `json:"ipv6Support,omitempty" tf:"ipv6_support,omitempty"`
	// +kubebuilder:validation:MinItems=1
	// +kubebuilder:validation:UniqueItems=true
	SubnetIDS []string `json:"subnetIDS" tf:"subnet_ids"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	TransitGatewayDefaultRouteTableAssociation bool `json:"transitGatewayDefaultRouteTableAssociation,omitempty" tf:"transit_gateway_default_route_table_association,omitempty"`
	// +optional
	TransitGatewayDefaultRouteTablePropagation bool   `json:"transitGatewayDefaultRouteTablePropagation,omitempty" tf:"transit_gateway_default_route_table_propagation,omitempty"`
	TransitGatewayID                           string `json:"transitGatewayID" tf:"transit_gateway_id"`
	VpcID                                      string `json:"vpcID" tf:"vpc_id"`
}

type Ec2TransitGatewayVpcAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// Ec2TransitGatewayVpcAttachmentList is a list of Ec2TransitGatewayVpcAttachments
type Ec2TransitGatewayVpcAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Ec2TransitGatewayVpcAttachment CRD objects
	Items []Ec2TransitGatewayVpcAttachment `json:"items,omitempty"`
}
