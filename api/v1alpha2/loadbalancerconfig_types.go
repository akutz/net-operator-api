// Copyright (c) 2020-2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// LoadBalancerConfigReady is added when the LoadBalancerConfig object has
	// been successfully realized.
	LoadBalancerConfigReady = "Ready"

	// LoadBalancerConfigFailure is added if any failure is encountered while
	// realizing LoadBalancerConfig object.
	LoadBalancerConfigFailure = "Failure"

	// LoadBalancerConfigIPPoolPressure condition status is set to True when
	// IPPool is low on free IPs.
	LoadBalancerConfigIPPoolPressure = "IPPoolPressure"
)

// LoadBalancerConfigSpec defines the desired state of LoadBalancerConfig
type LoadBalancerConfigSpec struct {
	// ProviderRef describes the resource that provides configuration details
	// for the load balancer.
	ProviderRef corev1.LocalObjectReference `json:"providerRef"`
}

// LoadBalancerConfigStatus defines the observed state of LoadBalancerConfig
type LoadBalancerConfigStatus struct {
	// +optional

	// Conditions describes the observed conditions of the resource.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// LoadBalancerConfig is the Schema for the LoadBalancerConfigs API
type LoadBalancerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadBalancerConfigSpec   `json:"spec,omitempty"`
	Status LoadBalancerConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerConfigList contains a list of LoadBalancerConfig
type LoadBalancerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancerConfig `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&LoadBalancerConfig{}, &LoadBalancerConfigList{})
}
