// Copyright (c) 2020-2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// HAProxyLoadBalancerConfigSpec defines the configuration for an
// HAProxyLoadBalancerConfig instance. The spec is used to configure the
// HAProxyLoadBalancer instance to correctly route traffic to services.
// This spec supports HAProxyLoadBalancerConfig Dataplane API 2.0+ sidecar
type HAProxyLoadBalancerConfigSpec struct {

	// +kubebuilder:validation:MinItems=1

	// Endpoints is a list of the addresses for the DataPlane API servers
	// used to configure HAProxy. One or more dataplane API endpoints are
	// possible with single-node and multi-node (active/passive) topologies.
	// A valid endpoint URL adheres to the following format:
	// SCHEME://<HOST>[:<PORT>]/<API_VERSION>, ex. https://127.0.0.1:443/v1.
	Endpoints []string `json:"endpoints"`

	// +optional

	// ServerName describes the value used to verify the server's peer
	// certificate. This defaults to the host parses from the endpoint.
	// certificates, but may be overridden to support a different value.
	ServerName string `json:"serverName,omitempty"`

	// CredentialSecretRef describes the Secret resource that contains the
	// credentials used to access the HAProxy dataplane API endpoint(s).
	// The secret may contain the following fields:
	//
	//     * certificateAuthorityData - PEM-encoded CA cert bundle
	//     * clientCertificateData    - PEM-encoded client cert
	//     * clientKeyData            - PEM-encoded client key
	//     * username                 - Used for basic auth. Defaults to client.
	//     * password                 - Used for basic auth. Defaults to cert.
	//
	// The following YAML is an example of one such Secret resource:
	//
	//     apiVersion: v1
	//     kind: Secret
	//     metadata:
	//       name: haproxy-lb-config
	//       namespace: vmware-system-netop
	//     data:
	//       certificateAuthorityData: Cg==
	//       clientCertificateData: Cg==
	//       clientKeyData: Cg==
	//       username: Y2xpZW50
	//       password: Y2VydA==
	CredentialSecretRef corev1.SecretReference `json:"credentialSecretRef"`
}

// HAProxyLoadBalancerConfigStatus defines the observed state of an
// HAProxyLoadBalancerConfig resource.
type HAProxyLoadBalancerConfigStatus struct {
}

// +genclient
// +genclient:nonNamespaced
// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Cluster
// +kubebuilder:storageversion
// +kubebuilder:subresource:status

// HAProxyLoadBalancerConfig is the Schema for the HAProxyLoadBalancerConfigs
// API.
type HAProxyLoadBalancerConfig struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   HAProxyLoadBalancerConfigSpec   `json:"spec,omitempty"`
	Status HAProxyLoadBalancerConfigStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// HAProxyLoadBalancerConfigList contains a list of HAProxyLoadBalancerConfig.
type HAProxyLoadBalancerConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []HAProxyLoadBalancerConfig `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&HAProxyLoadBalancerConfig{}, &HAProxyLoadBalancerConfigList{})
}
