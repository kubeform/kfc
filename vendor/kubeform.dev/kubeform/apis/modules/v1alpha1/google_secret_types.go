package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

type GoogleSecret struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GoogleSecretSpec   `json:"spec,omitempty"`
	Status            GoogleSecretStatus `json:"status,omitempty"`
}

type GoogleSecretSpec struct {
	// +optional
	SecretRef   *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`
	ProviderRef core.LocalObjectReference  `json:"providerRef" tf:"-"`
	// +optional
	Source string `json:"source" tf:"source"`

	// +optional
	// The list of application names that will store secrets
	ApplicationList []string `json:"applicationList,omitempty" tf:"application_list,omitempty"`
	// +optional
	// GCP credentials fils
	CredentialsFilePath json.RawMessage `json:"credentialsFilePath,omitempty" tf:"credentials_file_path,omitempty"`
	// +optional
	// The list of environments for secrets
	EnvList []string `json:"envList,omitempty" tf:"env_list,omitempty"`
	// +optional
	// The name of the project this applies to
	ProjectName json.RawMessage `json:"projectName,omitempty" tf:"project_name,omitempty"`
}

type GoogleSecretOutput struct {
	//
	AppBuckets string `json:"appBuckets" tf:"app-buckets"`
	//
	SharedBuckets string `json:"sharedBuckets" tf:"shared-buckets"`
}

type GoogleSecretStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GoogleSecretOutput `json:"output,omitempty"`
	// +optional
	State string `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GoogleSecretList is a list of GoogleSecrets
type GoogleSecretList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GoogleSecret CRD objects
	Items []GoogleSecret `json:"items,omitempty"`
}
