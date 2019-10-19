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

type DataLakeAnalyticsAccount struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataLakeAnalyticsAccountSpec   `json:"spec,omitempty"`
	Status            DataLakeAnalyticsAccountStatus `json:"status,omitempty"`
}

type DataLakeAnalyticsAccountSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	DefaultStoreAccountName string `json:"defaultStoreAccountName" tf:"default_store_account_name"`
	Location                string `json:"location" tf:"location"`
	Name                    string `json:"name" tf:"name"`
	ResourceGroupName       string `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	Tier string `json:"tier,omitempty" tf:"tier,omitempty"`
}

type DataLakeAnalyticsAccountStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataLakeAnalyticsAccountSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataLakeAnalyticsAccountList is a list of DataLakeAnalyticsAccounts
type DataLakeAnalyticsAccountList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataLakeAnalyticsAccount CRD objects
	Items []DataLakeAnalyticsAccount `json:"items,omitempty"`
}
