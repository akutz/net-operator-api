// Copyright (c) 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const (
	// DefaultNetworkLabelName is the name of the label applied to the Network
	// resource that is the default network in a namespace.
	DefaultNetworkLabelName = "netoperator.vmware.com/default-network"
)

// NetworkSpec defines the desired state of a Network resource.
type NetworkSpec struct {

	// +optional
	// +listType=set

	// Capabilities describes the capabilities of a given network, ex.
	// "multicast", "private", etc.
	// A capability must adhere to the format of a Kubernetes label value.
	Capabilities []string `json:"capabilities,omitempty"`

	// +optional

	// ControllerName describes the name of the controller responsible for
	// reconciling Network resources.
	//
	// When omitted, controllers reconciling Network resources determine
	// the default controller name from the environment variable
	// DEFAULT_NETWORK_CONTROLLER_NAME. If this environment variable is not
	// defined or empty, it defaults to netoperator.vmware.com/network.
	//
	// Once a non-empty value is assigned to this field, attempts to set this
	// field to an empty value will be silently ignored.
	ControllerName string `json:"controllerName,omitempty"`

	// +optional

	// DHCP4 describes whether or not the network uses DHCP for allocating IP4
	// addresses.
	DHCP4 bool `json:"dhcp4,omitempty"`

	// +optional

	// DHCP6 describes whether or not the network uses DHCP for allocating IP6
	// addresses.
	DHCP6 bool `json:"dhcp6,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=cidr

	// Gateway4 describes the default, IP4 gateway for this network.
	//
	// The address includes the network prefix length, ex. 192.168.0.1/24.
	Gateway4 string `json:"gateway4,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=cidr

	// Gateway6 describes the default, IP6 gateway for this interface.
	//
	// The address includes the network prefix length, ex. 2001:db8:101::1/64.
	Gateway6 string `json:"gateway6,omitempty"`

	// +optional
	// +listType=map
	// +listMapKey=clusterID

	// Identity describes the identify of the network for one or more given
	// vSphere clusters.
	Identity []NetworkIdentity `json:"identity,omitempty"`

	// +optional

	// ID describes the piece of information that is used to look up the
	// network object on the underlying infrastructure provider.
	//
	// Please note the meaning of this value is determined by the underlying
	// infrastructure provider.
	ID string `json:"id,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// IPAssignmentModes describes the supported modes of assigning IP addresses
	// on this network.
	IPAssignmentModes []IPAssignmentMode `json:"ipAssignmentModes,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// IPFamilies describes the IP families supported on this network.
	IPFamilies []IPFamily `json:"ipFamilies"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// IPRanges describes the IP4 and IP6 address ranges supported by this
	// network.
	//
	// Please note each IP range adheres to the CIDR format, ex. 192.168.0.0/24
	// or 2001:db8:101::0/64.
	//
	// Additionally, an IP range can be a single IP address, ex. 192.168.0.2/32
	// or 2001:db8:101::2/128.
	//
	// Please note if the network's IP assignment modes include DHCP or IPAM,
	// then any of the described ranges may be subject to those modes. A single
	// Network resource is not designed to represent multiple, smaller, logical
	// networks. The purpose for multiple ranges is to support multiple address
	// families as well as a single, non-contiguous, logical network.
	IPRanges []string `json:"ipRanges,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// Nameservers is a list of IP4 and/or IP6 addresses used as DNS
	// nameservers.
	//
	// Please note that Linux allows only three nameservers
	// (https://linux.die.net/man/5/resolv.conf).
	Nameservers []string `json:"nameservers,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// NTP describes a list of FQDNs, IP4, or IP6 addresses used as NTP servers.
	NTP []string `json:"ntp,omitempty"`

	// +optional

	// ProviderConfig describes additional configuration information for the
	// underlying network provider.
	ProviderConfig *ProviderConfig `json:"providerConfig,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// SearchDomains is a list of search domains used when resolving IP
	// addresses with DNS.
	SearchDomains []string `json:"searchDomains,omitempty"`
}

// NetworkStatus defines the observed state of a Network resource.
type NetworkStatus struct {

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// Capabilities describes the capabilities of a given network, ex.
	// "multicast", "private", etc.
	//
	// Every capability is also added to the resource's labels as
	// capability.netoperator.vmware.com/ + <capability>. For example, if the
	// capability is "multicast" then the following label will be added to the
	// resource: capability.netoperator.vmware.com/multicast.
	Capabilities []string `json:"capabilities,omitempty"`

	// +optional

	// Conditions describes the observed conditions of the resource.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=cidr

	// Gateway4 describes the default, IP4 gateway for this network.
	//
	// The address includes the network prefix length, ex. 192.168.0.1/24.
	Gateway4 string `json:"gateway4,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=cidr

	// Gateway6 describes the default, IP6 gateway for this interface.
	//
	// The address includes the network prefix length, ex. 2001:db8:101::1/64.
	Gateway6 string `json:"gateway6,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// IPAssignmentModes describes the supported modes of assigning IP addresses
	// on this network.
	IPAssignmentModes []IPAssignmentMode `json:"ipAssignmentModes,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// IPFamilies describes the IP families supported on this network.
	IPFamilies []IPFamily `json:"ipFamilies,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// IPRanges describes the IP4 and IP6 address ranges supported by this
	// network.
	//
	// Please note each IP range adheres to the CIDR format, ex. 192.168.0.0/24
	// or 2001:db8:101::0/64.
	//
	// Additionally, an IP range can be a single IP address, ex. 192.168.0.2/32
	// or 2001:db8:101::2/128.
	//
	// Please note if the network's IP assignment modes include DHCP or IPAM,
	// then any of the described ranges may be subject to those modes. A single
	// Network resource is not designed to represent multiple, smaller, logical
	// networks. The purpose for multiple ranges is to support multiple address
	// families as well as a single, non-contiguous, logical network.
	IPRanges []string `json:"ipRanges,omitempty"`

	// +optional

	// ProviderRef describes the Kubernetes, network-like resource that provides
	// this network's underlying capabilities.
	ProviderRef *corev1.TypedLocalObjectReference `json:"providerRef,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// Nameservers is a list of IP4 and/or IP6 addresses used as DNS
	// nameservers.
	//
	// Please note that Linux allows only three nameservers
	// (https://linux.die.net/man/5/resolv.conf).
	Nameservers []string `json:"nameservers,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// NTP describes a list of FQDNs, IP4, or IP6 addresses used as NTP servers.
	NTP []string `json:"ntp,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// SearchDomains is a list of search domains used when resolving IP
	// addresses with DNS.
	SearchDomains []string `json:"searchDomains,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=net
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="IPFamily",type="string",JSONPath=".status.ipFamily[*]"
// +kubebuilder:printcolumn:name="IPAssign",type="string",JSONPath=".status.ipAssign[*]"
// +kubebuilder:printcolumn:name="Gateway4",type="string",JSONPath=".status.gateway4"
// +kubebuilder:printcolumn:name="Gateway6",type="string",JSONPath=".status.gateway6"
// +kubebuilder:printcolumn:name="IPRanges",type="string",JSONPath=".status.ipRanges[*]"

// Network is the schema for the network API.
//
// A network resource surfaces an underlying, network infrastructure resource,
// enabling users to get and list information about all of a namespace's
// available networks without needing to know the underlying, network
// implementation.
//
// This type of network resource is akin to an IP subnetwork. While this
// resource may indicate multiple IP ranges, all of the ranges of a given
// address family (IP4 or IP6) share a single gateway.
type Network struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkSpec   `json:"spec,omitempty"`
	Status NetworkStatus `json:"status,omitempty"`
}

// NamespacedName returns this resource's namespace and name joined with a
// "/" character.
func (r Network) NamespacedName() string {
	return r.Namespace + "/" + r.Name
}

// +kubebuilder:object:root=true

// NetworkList contains a list of Network.
type NetworkList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Network `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&Network{}, &NetworkList{})
}
