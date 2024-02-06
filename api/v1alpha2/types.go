// Copyright (c) 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
)

// +kubebuilder:validation:Enum=dhcp4;dhcp6;ipam4;ipam6;none4;none6

// IPAssignmentMode defines the various methods for assigning IP addresses.
type IPAssignmentMode string

const (
	// IPAssignmentModeDHCP4 indicates DHCP is used for IP4 addresses.
	IPAssignmentModeDHCP4 IPAssignmentMode = "dhcp4"

	// IPAssignmentModeDHCP6 indicates DHCP is used for IP6 addresses.
	IPAssignmentModeDHCP6 IPAssignmentMode = "dhcp6"

	// IPAssignmentModeIPAM4 indicates IPAM is used for IP4 addresses.
	IPAssignmentModeIPAM4 IPAssignmentMode = "ipam4"

	// IPAssignmentModeIPAM6 indicates IPAM is used for IP6 addresses.
	IPAssignmentModeIPAM6 IPAssignmentMode = "ipam6"

	// IPAssignmentModeNone4 indicates no assignment is used for IP4 addresses.
	IPAssignmentModeNone4 IPAssignmentMode = "none4"

	// IPAssignmentModeNone6 indicates no assignment is used for IP6 addresses.
	IPAssignmentModeNone6 IPAssignmentMode = "none6"
)

// +kubebuilder:validation:Enum=ip4;ip6

// IPFamily defines the different IP address families.
type IPFamily string

const (
	// IPFamilyIP4 is the IP4 address family.
	IPFamilyIP4 IPFamily = "ip4"

	// IPFamilyIP6 is the IP6 address family.
	IPFamilyIP6 IPFamily = "ip6"
)

type NetworkIdentity struct {
	// ClusterID describes a value used to identify a vSphere cluster.
	ClusterID string `json:"clusterID"`

	// NetworkID describes the value used to identify the network resource for
	// the associated vSphere cluster.
	NetworkID string `json:"networkID"`
}

type KeyValuePair struct {
	Key string `json:"key"`

	// +optional

	Value string `json:"value,omitempty"`
}

type ProviderConfig struct {

	// +optional
	// +listType=map
	// +listMapKey=key

	// Properties describes additional configuration information for the
	// underlying network provider.
	Properties []KeyValuePair `json:"properties,omitempty"`

	// +optional

	// ProviderRef describes an object that contains additional configuration
	// for the underlying network provider.
	ProviderRef *corev1.TypedLocalObjectReference `json:"providerRef,omitempty"`
}
