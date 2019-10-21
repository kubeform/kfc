package v1alpha1

import (
	"encoding/json"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	base "kubeform.dev/kubeform/apis/base/v1alpha1"
)

// +genclient
// +k8s:openapi-gen=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Phase",type=string,JSONPath=`.status.phase`

type AutoscaleSetting struct {
	metav1.TypeMeta   `json:",inline,omitempty"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              AutoscaleSettingSpec   `json:"spec,omitempty"`
	Status            AutoscaleSettingStatus `json:"status,omitempty"`
}

type AutoscaleSettingSpecNotificationEmail struct {
	// +optional
	CustomEmails []string `json:"customEmails,omitempty" tf:"custom_emails,omitempty"`
	// +optional
	SendToSubscriptionAdministrator bool `json:"sendToSubscriptionAdministrator,omitempty" tf:"send_to_subscription_administrator,omitempty"`
	// +optional
	SendToSubscriptionCoAdministrator bool `json:"sendToSubscriptionCoAdministrator,omitempty" tf:"send_to_subscription_co_administrator,omitempty"`
}

type AutoscaleSettingSpecNotificationWebhook struct {
	// +optional
	Properties map[string]string `json:"properties,omitempty" tf:"properties,omitempty"`
	ServiceURI string            `json:"serviceURI" tf:"service_uri"`
}

type AutoscaleSettingSpecNotification struct {
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Email []AutoscaleSettingSpecNotificationEmail `json:"email,omitempty" tf:"email,omitempty"`
	// +optional
	Webhook []AutoscaleSettingSpecNotificationWebhook `json:"webhook,omitempty" tf:"webhook,omitempty"`
}

type AutoscaleSettingSpecProfileCapacity struct {
	Default int `json:"default" tf:"default"`
	Maximum int `json:"maximum" tf:"maximum"`
	Minimum int `json:"minimum" tf:"minimum"`
}

type AutoscaleSettingSpecProfileFixedDate struct {
	End   string `json:"end" tf:"end"`
	Start string `json:"start" tf:"start"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type AutoscaleSettingSpecProfileRecurrence struct {
	Days []string `json:"days" tf:"days"`
	// +kubebuilder:validation:MaxItems=1
	Hours []int64 `json:"hours" tf:"hours"`
	// +kubebuilder:validation:MaxItems=1
	Minutes []int64 `json:"minutes" tf:"minutes"`
	// +optional
	Timezone string `json:"timezone,omitempty" tf:"timezone,omitempty"`
}

type AutoscaleSettingSpecProfileRuleMetricTrigger struct {
	MetricName       string      `json:"metricName" tf:"metric_name"`
	MetricResourceID string      `json:"metricResourceID" tf:"metric_resource_id"`
	Operator         string      `json:"operator" tf:"operator"`
	Statistic        string      `json:"statistic" tf:"statistic"`
	Threshold        json.Number `json:"threshold" tf:"threshold"`
	TimeAggregation  string      `json:"timeAggregation" tf:"time_aggregation"`
	TimeGrain        string      `json:"timeGrain" tf:"time_grain"`
	TimeWindow       string      `json:"timeWindow" tf:"time_window"`
}

type AutoscaleSettingSpecProfileRuleScaleAction struct {
	Cooldown  string `json:"cooldown" tf:"cooldown"`
	Direction string `json:"direction" tf:"direction"`
	Type      string `json:"type" tf:"type"`
	Value     int    `json:"value" tf:"value"`
}

type AutoscaleSettingSpecProfileRule struct {
	// +kubebuilder:validation:MaxItems=1
	MetricTrigger []AutoscaleSettingSpecProfileRuleMetricTrigger `json:"metricTrigger" tf:"metric_trigger"`
	// +kubebuilder:validation:MaxItems=1
	ScaleAction []AutoscaleSettingSpecProfileRuleScaleAction `json:"scaleAction" tf:"scale_action"`
}

type AutoscaleSettingSpecProfile struct {
	// +kubebuilder:validation:MaxItems=1
	Capacity []AutoscaleSettingSpecProfileCapacity `json:"capacity" tf:"capacity"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	FixedDate []AutoscaleSettingSpecProfileFixedDate `json:"fixedDate,omitempty" tf:"fixed_date,omitempty"`
	Name      string                                 `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Recurrence []AutoscaleSettingSpecProfileRecurrence `json:"recurrence,omitempty" tf:"recurrence,omitempty"`
	// +optional
	// +kubebuilder:validation:MaxItems=10
	Rule []AutoscaleSettingSpecProfileRule `json:"rule,omitempty" tf:"rule,omitempty"`
}

type AutoscaleSettingSpec struct {
	ProviderRef core.LocalObjectReference `json:"providerRef" tf:"-"`

	ID string `json:"id,omitempty" tf:"id,omitempty"`

	// +optional
	Enabled  bool   `json:"enabled,omitempty" tf:"enabled,omitempty"`
	Location string `json:"location" tf:"location"`
	Name     string `json:"name" tf:"name"`
	// +optional
	// +kubebuilder:validation:MaxItems=1
	Notification []AutoscaleSettingSpecNotification `json:"notification,omitempty" tf:"notification,omitempty"`
	// +kubebuilder:validation:MaxItems=20
	Profile           []AutoscaleSettingSpecProfile `json:"profile" tf:"profile"`
	ResourceGroupName string                        `json:"resourceGroupName" tf:"resource_group_name"`
	// +optional
	Tags             map[string]string `json:"tags,omitempty" tf:"tags,omitempty"`
	TargetResourceID string            `json:"targetResourceID" tf:"target_resource_id"`
}

type AutoscaleSettingStatus struct {
	// Resource generation, which is updated on mutation by the API Server.
	// +optional
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
	// +optional
	Output *AutoscaleSettingSpec `json:"output,omitempty"`
	// +optional
	State *base.State `json:"state,omitempty"`
	// +optional
	Phase base.Phase `json:"phase,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +kubebuilder:object:root=true

// AutoscaleSettingList is a list of AutoscaleSettings
type AutoscaleSettingList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	// Items is a list of AutoscaleSetting CRD objects
	Items []AutoscaleSetting `json:"items,omitempty"`
}
