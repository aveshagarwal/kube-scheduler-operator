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
	// Fill me
	Kubeconfig string `json:"kubeconfig"`
	PolicyConfigFile string `json:"policyconfigfile"`
	KubeApiContentType string `json:"kubeapicontenttype"`
	KubeApiQps int32 `json:"kubeapiqps"`
	LeaderElect bool  `json:"leaderelect"`
	LeaderElectResourceLock string `json:"leaderelectresourcelock"`
	Port int32 `json:"port"`
}

type KubeSchedulerStatus struct {
	// Fill me
}
