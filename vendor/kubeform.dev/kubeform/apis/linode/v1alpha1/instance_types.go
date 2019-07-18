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

type Instance struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              InstanceSpec   `json:"spec,omitempty"`
	Status            InstanceStatus `json:"status,omitempty"`
}

type InstanceSpecConfig struct {
	// +optional
	Comments string `json:"comments,omitempty" tf:"comments,omitempty"`
	// +optional
	Kernel string `json:"kernel,omitempty" tf:"kernel,omitempty"`
	Label  string `json:"label" tf:"label"`
	// +optional
	MemoryLimit int `json:"memoryLimit,omitempty" tf:"memory_limit,omitempty"`
	// +optional
	RunLevel string `json:"runLevel,omitempty" tf:"run_level,omitempty"`
	// +optional
	VirtMode string `json:"virtMode,omitempty" tf:"virt_mode,omitempty"`
}

type InstanceSpecDisk struct {
	// +optional
	AuthorizedKeys []string `json:"authorizedKeys,omitempty" tf:"authorized_keys,omitempty"`
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty" tf:"authorized_users,omitempty"`
	Label           string   `json:"label" tf:"label"`
	// +optional
	RootPass string `json:"rootPass,omitempty" tf:"root_pass,omitempty"`
	Size     int    `json:"size" tf:"size"`
}

type InstanceSpec struct {
	// +optional
	AuthorizedKeys []string `json:"authorizedKeys,omitempty" tf:"authorized_keys,omitempty"`
	// +optional
	AuthorizedUsers []string `json:"authorizedUsers,omitempty" tf:"authorized_users,omitempty"`
	// +optional
	BackupID int `json:"backupID,omitempty" tf:"backup_id,omitempty"`
	// +optional
	Config []InstanceSpecConfig `json:"config,omitempty" tf:"config,omitempty"`
	// +optional
	Disk []InstanceSpecDisk `json:"disk,omitempty" tf:"disk,omitempty"`
	// +optional
	Group string `json:"group,omitempty" tf:"group,omitempty"`
	// +optional
	Image string `json:"image,omitempty" tf:"image,omitempty"`
	// +optional
	PrivateIP bool   `json:"privateIP,omitempty" tf:"private_ip,omitempty"`
	Region    string `json:"region" tf:"region"`
	// +optional
	RootPass string `json:"rootPass,omitempty" tf:"root_pass,omitempty"`
	// +optional
	StackscriptData map[string]string `json:"stackscriptData,omitempty" tf:"stackscript_data,omitempty"`
	// +optional
	StackscriptID int `json:"stackscriptID,omitempty" tf:"stackscript_id,omitempty"`
	// +optional
	// +kubebuilder:validation:UniqueItems=true
	Tags []string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Type string `json:"type,omitempty" tf:"type,omitempty"`
	// +optional
	WatchdogEnabled bool                      `json:"watchdogEnabled,omitempty" tf:"watchdog_enabled,omitempty"`
	ProviderRef     core.LocalObjectReference `json:"providerRef" tf:"-"`
}

type InstanceStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`

	TFState     []byte                `json:"tfState,omitempty"`
	TFStateHash string                `json:"tfStateHash,omitempty"`
	Output      *runtime.RawExtension `json:"output,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// InstanceList is a list of Instances
type InstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of Instance CRD objects
	Items []Instance `json:"items,omitempty"`
}
