package framework

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func deleteInForeground() *metav1.DeleteOptions {
	policy := metav1.DeletePropagationForeground
	return &metav1.DeleteOptions{PropagationPolicy: &policy}
}
