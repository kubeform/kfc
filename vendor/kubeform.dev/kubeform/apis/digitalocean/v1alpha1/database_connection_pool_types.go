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

type DatabaseConnectionPool struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              DatabaseConnectionPoolSpec   `json:"spec,omitempty"`
	Status            DatabaseConnectionPoolStatus `json:"status,omitempty"`
}

type DatabaseConnectionPoolSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	SecretRef *core.LocalObjectReference `json:"secretRef,omitempty" tf:"-"`

	ClusterID string `json:"clusterID" tf:"cluster_id"`
	DbName    string `json:"dbName" tf:"db_name"`
	// +optional
	Host string `json:"host,omitempty" tf:"host,omitempty"`
	Mode string `json:"mode" tf:"mode"`
	Name string `json:"name" tf:"name"`
	// +optional
	Password string `json:"-" sensitive:"true" tf:"password,omitempty"`
	// +optional
	Port int64 `json:"port,omitempty" tf:"port,omitempty"`
	// +optional
	PrivateHost string `json:"privateHost,omitempty" tf:"private_host,omitempty"`
	// +optional
	PrivateURI string `json:"-" sensitive:"true" tf:"private_uri,omitempty"`
	Size       int64  `json:"size" tf:"size"`
	// +optional
	Uri  string `json:"-" sensitive:"true" tf:"uri,omitempty"`
	User string `json:"user" tf:"user"`
}

type DatabaseConnectionPoolStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *DatabaseConnectionPoolSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// DatabaseConnectionPoolList is a list of DatabaseConnectionPools
type DatabaseConnectionPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of DatabaseConnectionPool CRD objects
	Items []DatabaseConnectionPool `json:"items,omitempty"`
}
