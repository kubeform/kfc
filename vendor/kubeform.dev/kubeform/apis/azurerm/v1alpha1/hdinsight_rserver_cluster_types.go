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

type HdinsightRserverCluster struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              HdinsightRserverClusterSpec   `json:"spec,omitempty"`
	Status            HdinsightRserverClusterStatus `json:"status,omitempty"`
}

type HdinsightRserverClusterSpecGateway struct {
	Enabled  bool   `json:"enabled" tf:"enabled"`
	Username string `json:"username" tf:"username"`
}

type HdinsightRserverClusterSpecRolesEdgeNode struct {
	// +optional
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightRserverClusterSpecRolesHeadNode struct {
	// +optional
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightRserverClusterSpecRolesWorkerNode struct {
	// +optional
	MinInstanceCount int `json:"minInstanceCount,omitempty" tf:"min_instance_count,omitempty"`
	// +optional
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID            string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	TargetInstanceCount int    `json:"targetInstanceCount" tf:"target_instance_count"`
	Username            string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightRserverClusterSpecRolesZookeeperNode struct {
	// +optional
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	SshKeys []string `json:"sshKeys,omitempty" tf:"ssh_keys,omitempty"`
	// +optional
	SubnetID string `json:"subnetID,omitempty" tf:"subnet_id,omitempty"`
	Username string `json:"username" tf:"username"`
	// +optional
	VirtualNetworkID string `json:"virtualNetworkID,omitempty" tf:"virtual_network_id,omitempty"`
	VmSize           string `json:"vmSize" tf:"vm_size"`
}

type HdinsightRserverClusterSpecRoles struct {
	// +kubebuilder:validation:MaxItems=1
	EdgeNode []HdinsightRserverClusterSpecRolesEdgeNode `json:"edgeNode" tf:"edge_node"`
	// +kubebuilder:validation:MaxItems=1
	HeadNode []HdinsightRserverClusterSpecRolesHeadNode `json:"headNode" tf:"head_node"`
	// +kubebuilder:validation:MaxItems=1
	WorkerNode []HdinsightRserverClusterSpecRolesWorkerNode `json:"workerNode" tf:"worker_node"`
	// +kubebuilder:validation:MaxItems=1
	ZookeeperNode []HdinsightRserverClusterSpecRolesZookeeperNode `json:"zookeeperNode" tf:"zookeeper_node"`
}

type HdinsightRserverClusterSpecStorageAccount struct {
	IsDefault          bool   `json:"isDefault" tf:"is_default"`
	StorageContainerID string `json:"storageContainerID" tf:"storage_container_id"`
}

type HdinsightRserverClusterSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	Secret *core.LocalObjectReference `json:"secret,omitempty" tf:"-"`

	ClusterVersion string `json:"clusterVersion" tf:"cluster_version"`
	// +kubebuilder:validation:MaxItems=1
	Gateway           []HdinsightRserverClusterSpecGateway `json:"gateway" tf:"gateway"`
	Location          string                               `json:"location" tf:"location"`
	Name              string                               `json:"name" tf:"name"`
	ResourceGroupName string                               `json:"resourceGroupName" tf:"resource_group_name"`
	// +kubebuilder:validation:MaxItems=1
	Roles          []HdinsightRserverClusterSpecRoles          `json:"roles" tf:"roles"`
	Rstudio        bool                                        `json:"rstudio" tf:"rstudio"`
	StorageAccount []HdinsightRserverClusterSpecStorageAccount `json:"storageAccount" tf:"storage_account"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	Tier string            `json:"tier" tf:"tier"`
}

type HdinsightRserverClusterStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// HdinsightRserverClusterList is a list of HdinsightRserverClusters
type HdinsightRserverClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of HdinsightRserverCluster CRD objects
	Items []HdinsightRserverCluster `json:"items,omitempty"`
}