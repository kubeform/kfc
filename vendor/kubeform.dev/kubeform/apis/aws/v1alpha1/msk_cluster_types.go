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

type MskCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              MskClusterSpec   `json:"spec,omitempty"`
	Status            MskClusterStatus `json:"status,omitempty"`
}

type MskClusterSpecBrokerNodeGroupInfo struct {
	// +optional
	AzDistribution string   `json:"azDistribution,omitempty" tf:"az_distribution,omitempty"`
	ClientSubnets  []string `json:"clientSubnets" tf:"client_subnets"`
	EbsVolumeSize  int64    `json:"ebsVolumeSize" tf:"ebs_volume_size"`
	InstanceType   string   `json:"instanceType" tf:"instance_type"`
	SecurityGroups []string `json:"securityGroups" tf:"security_groups"`
}

type MskClusterSpecClientAuthenticationTls struct {
	// +optional
	CertificateAuthorityArns []string `json:"certificateAuthorityArns,omitempty" tf:"certificate_authority_arns,omitempty"`
}

type MskClusterSpecClientAuthentication struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Tls []MskClusterSpecClientAuthenticationTls `json:"tls,omitempty" tf:"tls,omitempty"`
}

type MskClusterSpecConfigurationInfo struct {
	Arn      string `json:"arn" tf:"arn"`
	Revision int64  `json:"revision" tf:"revision"`
}

type MskClusterSpecEncryptionInfoEncryptionInTransit struct {
	// +optional
	ClientBroker string `json:"clientBroker,omitempty" tf:"client_broker,omitempty"`
	// +optional
	InCluster bool `json:"inCluster,omitempty" tf:"in_cluster,omitempty"`
}

type MskClusterSpecEncryptionInfo struct {
	// +optional
	EncryptionAtRestKmsKeyArn string `json:"encryptionAtRestKmsKeyArn,omitempty" tf:"encryption_at_rest_kms_key_arn,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionInTransit []MskClusterSpecEncryptionInfoEncryptionInTransit `json:"encryptionInTransit,omitempty" tf:"encryption_in_transit,omitempty"`
}

type MskClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Arn string `json:"arn,omitempty" tf:"arn,omitempty"`
	// +optional
	BootstrapBrokers string `json:"bootstrapBrokers,omitempty" tf:"bootstrap_brokers,omitempty"`
	// +optional
	BootstrapBrokersTLS string `json:"bootstrapBrokersTLS,omitempty" tf:"bootstrap_brokers_tls,omitempty"`
	// +kubebuilder:validation:MaxItems=1
	BrokerNodeGroupInfo []MskClusterSpecBrokerNodeGroupInfo `json:"brokerNodeGroupInfo" tf:"broker_node_group_info"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ClientAuthentication []MskClusterSpecClientAuthentication `json:"clientAuthentication,omitempty" tf:"client_authentication,omitempty"`
	ClusterName          string                               `json:"clusterName" tf:"cluster_name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	ConfigurationInfo []MskClusterSpecConfigurationInfo `json:"configurationInfo,omitempty" tf:"configuration_info,omitempty"`
	// +optional
	CurrentVersion string `json:"currentVersion,omitempty" tf:"current_version,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	EncryptionInfo []MskClusterSpecEncryptionInfo `json:"encryptionInfo,omitempty" tf:"encryption_info,omitempty"`
	// +optional
	EnhancedMonitoring  string `json:"enhancedMonitoring,omitempty" tf:"enhanced_monitoring,omitempty"`
	KafkaVersion        string `json:"kafkaVersion" tf:"kafka_version"`
	NumberOfBrokerNodes int64  `json:"numberOfBrokerNodes" tf:"number_of_broker_nodes"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	ZookeeperConnectString string `json:"zookeeperConnectString,omitempty" tf:"zookeeper_connect_string,omitempty"`
}

type MskClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *MskClusterSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// MskClusterList is a list of MskClusters
type MskClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of MskCluster CRD objects
	Items []MskCluster `json:"items,omitempty"`
}
