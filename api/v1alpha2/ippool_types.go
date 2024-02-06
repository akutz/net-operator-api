// Copyright (c) 2020-2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IPAMDisabledAnnotationKeyName is the name of the annotation added to
// GatewayClass resources that do not participate in net-operator's IPAM.
// The value does not need to be truthy; the presence of the key is what
// disables net-operator's IPAM for that GatewayClass.
const IPAMDisabledAnnotationKeyName = "netoperator.vmware.com/ipam-disabled"

const (
	// IPPoolFull condition is added when no more IPs are free in the pool.
	IPPoolFull = "full"

	// IPPoolReady condition is added when IPPool has been realized.
	IPPoolReady = "ready"

	// IPPoolFail condition is added when an error was encountered in realizing.
	IPPoolFail = "failure"
)

// IPPoolSpec defines the desired state of IPPool
type IPPoolSpec struct {
	// StartingAddress represents the starting IP address of the pool.
	StartingAddress string `json:"startingAddress"`

	// AddressCount represents the number of IP addresses in the pool.
	AddressCount int64 `json:"addressCount"`
}

// IPPoolStatus defines the current state of IPPool.
type IPPoolStatus struct {
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

// IPPool is the Schema for the ippools API.
// It represents a pool of IP addresses that are owned and managed by the IPPool
// controller. Provider specific networks can associate themselves with IPPool
// objects to use network operator's IPAM implementation.
type IPPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IPPoolSpec   `json:"spec,omitempty"`
	Status IPPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// IPPoolList contains a list of IPPool
type IPPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPPool `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&IPPool{}, &IPPoolList{})
}
