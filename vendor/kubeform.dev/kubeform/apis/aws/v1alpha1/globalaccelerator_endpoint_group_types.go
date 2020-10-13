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

type GlobalacceleratorEndpointGroup struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              GlobalacceleratorEndpointGroupSpec   `json:"spec,omitempty"`
	Status            GlobalacceleratorEndpointGroupStatus `json:"status,omitempty"`
}

type GlobalacceleratorEndpointGroupSpecEndpointConfiguration struct {
	// +optional
	EndpointID string `json:"endpointID,omitempty" tf:"endpoint_id,omitempty"`
	// +optional
	Weight int64 `json:"weight,omitempty" tf:"weight,omitempty"`
}

type GlobalacceleratorEndpointGroupSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	RemoteBackend *base.Backend `json:"remoteBackend,omitempty" tf:"-"`

	// +optional
	// +kubebuilder:validation:MaxItems=10
	EndpointConfiguration []GlobalacceleratorEndpointGroupSpecEndpointConfiguration `json:"endpointConfiguration,omitempty" tf:"endpoint_configuration,omitempty"`
	// +optional
	EndpointGroupRegion string `json:"endpointGroupRegion,omitempty" tf:"endpoint_group_region,omitempty"`
	// +optional
	HealthCheckIntervalSeconds int64 `json:"healthCheckIntervalSeconds,omitempty" tf:"health_check_interval_seconds,omitempty"`
	// +optional
	HealthCheckPath string `json:"healthCheckPath,omitempty" tf:"health_check_path,omitempty"`
	// +optional
	HealthCheckPort int64 `json:"healthCheckPort,omitempty" tf:"health_check_port,omitempty"`
	// +optional
	HealthCheckProtocol string `json:"healthCheckProtocol,omitempty" tf:"health_check_protocol,omitempty"`
	ListenerArn         string `json:"listenerArn" tf:"listener_arn"`
	// +optional
	ThresholdCount int64 `json:"thresholdCount,omitempty" tf:"threshold_count,omitempty"`
	// +optional
	TrafficDialPercentage float64 `json:"trafficDialPercentage,omitempty" tf:"traffic_dial_percentage,omitempty"`
}

type GlobalacceleratorEndpointGroupStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *GlobalacceleratorEndpointGroupSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// GlobalacceleratorEndpointGroupList is a list of GlobalacceleratorEndpointGroups
type GlobalacceleratorEndpointGroupList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of GlobalacceleratorEndpointGroup CRD objects
	Items []GlobalacceleratorEndpointGroup `json:"items,omitempty"`
}
