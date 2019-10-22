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

type DataFactory struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DataFactorySpec   `json:"spec,omitempty"`
	Status            DataFactoryStatus `json:"status,omitempty"`
}

type DataFactorySpecGithubConfiguration struct {
	AccountName    string `json:"accountName" tf:"account_name"`
	BranchName     string `json:"branchName" tf:"branch_name"`
	GitURL         string `json:"gitURL" tf:"git_url"`
	RepositoryName string `json:"repositoryName" tf:"repository_name"`
	RootFolder     string `json:"rootFolder" tf:"root_folder"`
}

type DataFactorySpecIdentity struct {
	// +optional
	PrincipalID string `json:"principalID,omitempty" tf:"principal_id,omitempty"`
	// +optional
	TenantID string `json:"tenantID,omitempty" tf:"tenant_id,omitempty"`
	Type     string `json:"type" tf:"type"`
}

type DataFactorySpecVstsConfiguration struct {
	AccountName    string `json:"accountName" tf:"account_name"`
	BranchName     string `json:"branchName" tf:"branch_name"`
	ProjectName    string `json:"projectName" tf:"project_name"`
	RepositoryName string `json:"repositoryName" tf:"repository_name"`
	RootFolder     string `json:"rootFolder" tf:"root_folder"`
	TenantID       string `json:"tenantID" tf:"tenant_id"`
}

type DataFactorySpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:MaxItems=1
	GithubConfiguration []DataFactorySpecGithubConfiguration `json:"githubConfiguration,omitempty" tf:"github_configuration,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Identity          []DataFactorySpecIdentity `json:"identity,omitempty" tf:"identity,omitempty"`
	Location          string                    `json:"location" tf:"location"`
	Name              string                    `json:"name" tf:"name"`
	ResourceGroupName string                    `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	VstsConfiguration []DataFactorySpecVstsConfiguration `json:"vstsConfiguration,omitempty" tf:"vsts_configuration,omitempty"`
}

type DataFactoryStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DataFactorySpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DataFactoryList is a list of DataFactorys
type DataFactoryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DataFactory CRD objects
	Items []DataFactory `json:"items,omitempty"`
}
