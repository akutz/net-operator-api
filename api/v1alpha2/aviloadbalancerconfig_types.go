// Copyright (c) 2020-2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:validation:Enum=INFO;DEBUG;WARN;ERROR

// AviLoadBalancerLogLevel is a valid log level for the Avi Kubernetes Operator.
type AviLoadBalancerLogLevel string

const (
	// AviLoadBalancerLogLevelInfo is the INFO log level for AKO.
	AviLoadBalancerLogLevelInfo AviLoadBalancerLogLevel = "INFO"

	// AviLoadBalancerLogLevelDebug is the DEBUG log level for AKO.
	AviLoadBalancerLogLevelDebug AviLoadBalancerLogLevel = "DEBUG"

	// AviLoadBalancerLogLevelWarn is the WARN log level for AKO.
	AviLoadBalancerLogLevelWarn AviLoadBalancerLogLevel = "WARN"

	// AviLoadBalancerLogLevelError is the ERROR log level for AKO.
	AviLoadBalancerLogLevelError AviLoadBalancerLogLevel = "ERROR"
)

// +kubebuilder:validation:Enum=controller;supervisor

// AviLoadBalancerIPAMType is the type of IPAM used by Avi.
type AviLoadBalancerIPAMType string

const (
	// AviLoadBalancerSupervisorIPAM indicates that IPAM is provided by the
	// Supervisor cluster.
	AviLoadBalancerSupervisorIPAM AviLoadBalancerIPAMType = "supervisor"

	// AviLoadBalancerControllerIPAM indicates that IPAM is provided by the Avi
	// Controller.
	AviLoadBalancerControllerIPAM AviLoadBalancerIPAMType = "controller"
)

// AviLoadBalancerConfigSpec defines the configuration for an Avi load balancer.
// This specification is used to configure the resources the Avi Kubernetes
// Operator (AKO) requires in order to connect to the Avi load balancer.
type AviLoadBalancerConfigSpec struct {
	// Server is the endpoint for the AVI Controller REST API.
	// A valid endpoint URL adheres to the following format:
	// SCHEME://<HOST>[:<PORT>], ex. https://127.0.0.1:443.
	Server string `json:"server"`

	// +optional
	// +kubebuilder:default:=Default-Cloud

	// CloudName is used by the Avi Kubernetes Operator (AKO) when querying
	// properties via the Avi REST API, ex. /api/cloud/?name=CLOUD_NAME.
	// Defaults to Default-Cloud.
	CloudName string `json:"cloudName,omitempty"`

	// +optional
	// +kubebuilder:default:=true

	// AdvancedL4 is a flag that enables support for WCP in AKO.
	// Defaults to true.
	AdvancedL4 *bool `json:"advancedL4,omitempty"`

	// +optional
	// +kubebuilder:default:=WARN

	// LogLevel specifies the log level used by AKO.
	LogLevel AviLoadBalancerLogLevel `json:"logLevel,omitempty"`

	// +optional
	// +kubebuilder:default:=controller

	// IPAMType is the type of IPAM used by the Avi Software Load Balancer.
	IPAMType AviLoadBalancerIPAMType `json:"ipamType,omitempty"`

	// CredentialSecretRef describes the Secret resource that contains the
	// credentials used to access the HAProxy dataplane API endpoint(s).
	// The secret may contain the following fields:
	//
	//     * certificateAuthorityData - PEM-encoded CA cert bundle
	//     * username                 - Used for basic auth.
	//     * password                 - Used for basic auth.
	//
	// The following YAML is an example of one such Secret resource:
	//
	//     apiVersion: v1
	//     kind: Secret
	//     metadata:
	//       name: avi-lb-config
	//       namespace: vmware-system-netop
	//     data:
	//       certificateAuthorityData: Cg==
	//       username: Cg==
	//       password: Cg==
	CredentialSecretRef corev1.SecretReference `json:"credentialSecretRef"`
}

// AviLoadBalancerConfigStatus is unused because AviLoadBalancerConfigSpec is
// purely a configuration resource.
type AviLoadBalancerConfigStatus struct {
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// AviLoadBalancerConfig is the Schema for the AviLoadBalancerConfigs API
type AviLoadBalancerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AviLoadBalancerConfigSpec   `json:"spec,omitempty"`
	Status AviLoadBalancerConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AviLoadBalancerConfigList contains a list of AviLoadBalancerConfig
type AviLoadBalancerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []AviLoadBalancerConfig `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&AviLoadBalancerConfig{}, &AviLoadBalancerConfigList{})
}
