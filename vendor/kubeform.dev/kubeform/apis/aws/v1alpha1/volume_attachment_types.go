package v1alpha1

import (
	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type VolumeAttachment struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              VolumeAttachmentSpec   `json:"spec,omitempty"`
	Status            VolumeAttachmentStatus `json:"status,omitempty"`
}

type VolumeAttachmentSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	DeviceName string `json:"deviceName" tf:"device_name"`
	// +optional
	ForceDetach bool   `json:"forceDetach,omitempty" tf:"force_detach,omitempty"`
	InstanceID  string `json:"instanceID" tf:"instance_id"`
	// +optional
	SkipDestroy bool   `json:"skipDestroy,omitempty" tf:"skip_destroy,omitempty"`
	VolumeID    string `json:"volumeID" tf:"volume_id"`
}

type VolumeAttachmentStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
	TFState            string `json:"tfState,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// VolumeAttachmentList is a list of VolumeAttachments
type VolumeAttachmentList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of VolumeAttachment CRD objects
	Items []VolumeAttachment `json:"items,omitempty"`
}
