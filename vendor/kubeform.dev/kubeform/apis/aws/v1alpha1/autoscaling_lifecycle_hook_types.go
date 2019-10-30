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

type AutoscalingLifecycleHook struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscalingLifecycleHookSpec   `json:"spec,omitempty"`
	Status            AutoscalingLifecycleHookStatus `json:"status,omitempty"`
}

type AutoscalingLifecycleHookSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	AutoscalingGroupName string `json:"autoscalingGroupName" tf:"autoscaling_group_name"`
	// +optional
	DefaultResult string `json:"defaultResult,omitempty" tf:"default_result,omitempty"`
	// +optional
	HeartbeatTimeout    int    `json:"heartbeatTimeout,omitempty" tf:"heartbeat_timeout,omitempty"`
	LifecycleTransition string `json:"lifecycleTransition" tf:"lifecycle_transition"`
	Name                string `json:"name" tf:"name"`
	// +optional
	NotificationMetadata string `json:"notificationMetadata,omitempty" tf:"notification_metadata,omitempty"`
	// +optional
	NotificationTargetArn string `json:"notificationTargetArn,omitempty" tf:"notification_target_arn,omitempty"`
	// +optional
	RoleArn string `json:"roleArn,omitempty" tf:"role_arn,omitempty"`
}

type AutoscalingLifecycleHookStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutoscalingLifecycleHookSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscalingLifecycleHookList is a list of AutoscalingLifecycleHooks
type AutoscalingLifecycleHookList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscalingLifecycleHook CRD objects
	Items []AutoscalingLifecycleHook `json:"items,omitempty"`
}
