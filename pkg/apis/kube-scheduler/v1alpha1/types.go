package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KubeSchedulerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`
	Items           []KubeScheduler `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type KubeScheduler struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata"`
	Spec              KubeSchedulerSpec   `json:"spec"`
	Status            KubeSchedulerStatus `json:"status,omitempty"`
}

type KubeSchedulerSpec struct {
	ConfigFileConfigMapName string            `json:"ConfigFileConfigMapName"`
	NodeSelector            map[string]string `json:"nodeSelector,omitempty"`
}

type KubeSchedulerStatus struct {
	// Fill me
}
