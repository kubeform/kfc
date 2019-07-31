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

type RedisInstance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              RedisInstanceSpec   `json:"spec,omitempty"`
	Status            RedisInstanceStatus `json:"status,omitempty"`
}

type RedisInstanceSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	// +optional
	AlternativeLocationID string `json:"alternativeLocationID,omitempty" tf:"alternative_location_id,omitempty"`
	// +optional
	AuthorizedNetwork string `json:"authorizedNetwork,omitempty" tf:"authorized_network,omitempty"`
	// +optional
	DisplayName string `json:"displayName,omitempty" tf:"display_name,omitempty"`
	// +optional
	Labels map[string]string `json:"labels,omitempty" tf:"labels,omitempty"`
	// +optional
	LocationID   string `json:"locationID,omitempty" tf:"location_id,omitempty"`
	MemorySizeGb int    `json:"memorySizeGb" tf:"memory_size_gb"`
	Name         string `json:"name" tf:"name"`
	// +optional
	Project string `json:"project,omitempty" tf:"project,omitempty"`
	// +optional
	RedisConfigs map[string]string `json:"redisConfigs,omitempty" tf:"redis_configs,omitempty"`
	// +optional
	RedisVersion string `json:"redisVersion,omitempty" tf:"redis_version,omitempty"`
	// +optional
	Region string `json:"region,omitempty" tf:"region,omitempty"`
	// +optional
	ReservedIPRange string `json:"reservedIPRange,omitempty" tf:"reserved_ip_range,omitempty"`
	// +optional
	Tier string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type RedisInstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState *runtime.RawExtension `json:"tfState,omitempty"`
	Output  *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// RedisInstanceList is a list of RedisInstances
type RedisInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of RedisInstance CRD objects
	Items []RedisInstance `json:"items,omitempty"`
}