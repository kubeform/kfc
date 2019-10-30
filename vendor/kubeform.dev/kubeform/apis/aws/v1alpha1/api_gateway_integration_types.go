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

type ApiGatewayIntegration struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApiGatewayIntegrationSpec   `json:"spec,omitempty"`
	Status            ApiGatewayIntegrationStatus `json:"status,omitempty"`
}

type ApiGatewayIntegrationSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	CacheKeyParameters []string `json:"cacheKeyParameters,omitempty" tf:"cache_key_parameters,omitempty"`
	// +optional
	CacheNamespace string `json:"cacheNamespace,omitempty" tf:"cache_namespace,omitempty"`
	// +optional
	ConnectionID string `json:"connectionID,omitempty" tf:"connection_id,omitempty"`
	// +optional
	ConnectionType string `json:"connectionType,omitempty" tf:"connection_type,omitempty"`
	// +optional
	ContentHandling string `json:"contentHandling,omitempty" tf:"content_handling,omitempty"`
	// +optional
	Credentials string `json:"credentials,omitempty" tf:"credentials,omitempty"`
	HttpMethod  string `json:"httpMethod" tf:"http_method"`
	// +optional
	IntegrationHTTPMethod string `json:"integrationHTTPMethod,omitempty" tf:"integration_http_method,omitempty"`
	// +optional
	PassthroughBehavior string `json:"passthroughBehavior,omitempty" tf:"passthrough_behavior,omitempty"`
	// +optional
	RequestParameters map[string]string `json:"requestParameters,omitempty" tf:"request_parameters,omitempty"`
	// +optional
	RequestTemplates map[string]string `json:"requestTemplates,omitempty" tf:"request_templates,omitempty"`
	ResourceID       string            `json:"resourceID" tf:"resource_id"`
	RestAPIID        string            `json:"restAPIID" tf:"rest_api_id"`
	// +optional
	TimeoutMilliseconds int    `json:"timeoutMilliseconds,omitempty" tf:"timeout_milliseconds,omitempty"`
	Type                string `json:"type" tf:"type"`
	// +optional
	Uri string `json:"uri,omitempty" tf:"uri,omitempty"`
}

type ApiGatewayIntegrationStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *ApiGatewayIntegrationSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// ApiGatewayIntegrationList is a list of ApiGatewayIntegrations
type ApiGatewayIntegrationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of ApiGatewayIntegration CRD objects
	Items []ApiGatewayIntegration `json:"items,omitempty"`
}
