// Copyright (c) 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NetworkInterfacePortAllocation describes the desired port allocation.
type NetworkInterfacePortAllocation struct {
	// NodeName describes the name of the node where the port must be allocated.
	NodeName string `json:"nodeName"`
}

// RouteSpec defines a static route for a guest.
type RouteSpec struct {
	// To is an IP4 address.
	To string `json:"to"`

	// Via is an IP4 address.
	Via string `json:"via"`

	// Metric is the weight/priority of the route.
	Metric int32 `json:"metric"`
}

// NetworkInterface describes the desired state of a networkinterfaces resource.
type NetworkInterfaceSpec struct {
	// +optional

	// AttachedTo specifies the resource to which this interface is attached.
	//
	// Please note specified this value does not cause the interface to be
	// attached to the resource. This field enables consumers of this resource
	// to indicate to which resource this interface has been attached.
	// Similarly, consumers should ensure this field is set to its empty value
	// when the interface is no longer attached.
	AttachedTo *corev1.TypedLocalObjectReference `json:"attachedTo,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// Addresses is an optional list of IP4 or IP6 addresses to assign to the
	// interface.
	//
	// Please note this field is only supported if the network supports manual
	// IP allocation.
	//
	// Please note IP4 and IP6 addresses must include the network prefix length,
	// ex. 192.168.0.10/24 or 2001:db8:101::a/64.
	//
	// Please note this field may not contain IP4 addresses if DHCP4 is set
	// to true or IP6 addresses if DHCP6 is set to true.
	//
	// Please note this field may be used concurrently with the AddressNames
	// property.
	Addresses []string `json:"addresses,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// AddressNames describes a list of names of IPAddress resources used as the
	// source of IP addresses for this resource. The named IPAddress resources
	// must be in the same namespace as this resource.
	//
	// Please note this field may be used concurrently with the Addresses
	// property.
	AddressNames []string `json:"addressNames,omitempty"`

	// +optional

	// DHCP4 indicates whether or not to use DHCP for IP4 networking.
	//
	// Please note this field is only supported if the network supports DHCP4.
	//
	// Please note this field is mutually exclusive with IP4 addresses in the
	// Addresses field and the Gateway4 field.
	DHCP4 bool `json:"dhcp4,omitempty"`

	// +optional

	// DHCP6 indicates whether or not to use DHCP for IP6 networking.
	//
	// Please note this field is only supported if the network supports DHCP6.
	//
	// Please note this field is mutually exclusive with IP4 addresses in the
	// Addresses field and the Gateway6 field.
	DHCP6 bool `json:"dhcp6,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=cidr

	// Gateway4 is the default, IP4 gateway for this interface.
	//
	// Please note this field is only supported if the network supports manual
	// IP allocation.
	//
	// Please note the IP address must include the network prefix length, ex.
	// 192.168.0.1/24.
	//
	// Most guests do not handle multiple network interfaces with default
	// gateway addresses, so clients using this resource should take care to
	// only connect a single interface with a non-empty value for this field.
	//
	// Please note this field is mutually exclusive with DHCP4.
	Gateway4 string `json:"gateway4,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=cidr

	// Gateway6 is the default, IP6 gateway for this interface.
	//
	// Please note this field is only supported if the network supports manual
	// IP allocation.
	//
	// Please note the IP address must include the network prefix length, ex.
	// 2001:db8:101::1/64.
	//
	// Most guests do not handle multiple network interfaces with default
	// gateway addresses, so clients using this resource should take care to
	// only connect a single interface with a non-empty value for this field.
	//
	// Please note this field is mutually exclusive with DHCP6.
	Gateway6 string `json:"gateway6,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=mac

	// MACAddr describes the desired MAC address for this network interface.
	MACAddr string `json:"macAddr,omitempty"`

	// +optional

	// MTU is the Maximum Transmission Unit size in bytes.
	//
	// Please note support for this field is dependent upon the bootstrap
	// provider used by the guest to which this interface is connected.
	MTU *int64 `json:"mtu,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// Nameservers is a list of IP4 and/or IP6 addresses used as DNS
	// nameservers.
	//
	// Please note support for this field is dependent upon the bootstrap
	// provider used by the guest to which this interface is connected.
	//
	// Please note that Linux allows only three nameservers
	// (https://linux.die.net/man/5/resolv.conf).
	Nameservers []string `json:"nameservers,omitempty"`

	// +optional

	// NetworkName describes the name of the Network resource to which to
	// connect this interface.
	//
	// If omitted, this field defaults to the name of the namespace's default
	// network. The default network for a namespace is either the sole Network
	// resource in the namespace, or if multiple Network resources, the one with
	// the label netoperator.vmware.com/default-network.
	NetworkName string `json:"networkName,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// NTP describes a list of FQDNs, IP4, or IP6 addresses used as NTP servers.
	NTP []string `json:"ntp,omitempty"`

	// +optional

	// PortAllocation describes the desired port allocation behavior.
	// This useful when the network interface is attached to the network without
	// vCenter's involvement.
	PortAllocation *NetworkInterfacePortAllocation `json:"portAllocation,omitempty"`

	// +optional

	// ProviderConfig describes additional configuration information for the
	// underlying network provider.
	ProviderConfig *ProviderConfig `json:"providerConfig,omitempty"`

	// +optional

	// Routes is a list of optional, static routes.
	//
	// Please note support for this field is dependent upon the bootstrap
	// provider used by the guest to which this interface is connected.
	Routes []RouteSpec `json:"routes,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// SearchDomains is a list of search domains used when resolving IP
	// addresses with DNS.
	//
	// Please note support for this field is dependent upon the bootstrap
	// provider used by the guest to which this interface is connected.
	SearchDomains []string `json:"searchDomains,omitempty"`
}

// NetworkInterfaceStatus defines the observed state of a networkinterfaces
// resource.
type NetworkInterfaceStatus struct {
	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// Addresses describes the IP4 and/or IP6 addresses assigned to this network
	// interface.
	//
	// Please note IP4 and IP6 addresses will include the network prefix length,
	// ex. 192.168.0.10/24 or 2001:db8:101::a/64.
	//
	// This value will not contain any IP4 addresses when DHCP4 is true or IP6
	// addresses when DHCP6 is true.
	Addresses []string `json:"addresses,omitempty"`

	// +optional

	// Conditions describes the observed conditions of the resource.
	Conditions []metav1.Condition `json:"conditions,omitempty"`

	// +optional

	// ConnectionID identifies the connected port on the backing network, is
	// only valid on the requested node, and only set if port allocation was
	// requested.
	ConnectionID string `json:"connectionID,omitempty"`

	// +optional

	// DHCP4 describes whether or not the network to which this interface is
	// connected uses DHCP for allocating IP4 addresses.
	DHCP4 bool `json:"dhcp4,omitempty"`

	// +optional

	// DHCP6 describes whether or not the network to which this interface is
	// connected uses DHCP for allocating IP6 addresses.
	DHCP6 bool `json:"dhcp6,omitempty"`

	// +optional

	// ExternalID describes the external ID to assign to the VirtualEthernetCard
	// device.
	//
	// If non-empty, this value *must* be used as-is when configuring the
	// VirtualEthernetCard device, or else it may not function correctly.
	ExternalID string `json:"string,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=cidr

	// Gateway4 describes the default, IP4 gateway for the network to which this
	// network interface is connected.
	//
	// Please note address includes the network prefix length, ex.
	// 192.168.0.1/24.
	//
	// This value is omitted when DHCP4 is true.
	Gateway4 string `json:"gateway4,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=cidr

	// Gateway6 describes the default, IP6 gateway for the network to which this
	// network interface is connected.
	//
	// Please note address includes the network prefix length, ex.
	// 2001:db8:101::1/64.
	//
	// This value is omitted when DHCP6 is true.
	Gateway6 string `json:"gateway6,omitempty"`

	// +optional
	// +listType=map
	// +listMapKey=clusterID

	// Identity describes the identify of the network interface on a given
	// vSphere cluster.
	Identity []NetworkIdentity `json:"identity,omitempty"`

	// +optional
	// +kubebuilder:validation:Format:=mac

	// MACAddr describes the MAC address for this network interface.
	//
	// If non-empty, this value *must* be used as-is when configuring the
	// VirtualEthernetCard device, or else it may not function correctly.
	MACAddr string `json:"macAddr,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// Nameservers describes an optional list of IP4 and/or IP6 addresses
	// configured as DNS servers for this interface.
	//
	// This value will not contain any IP4 addresses when DHCP4 is true or
	// IP6 addresses when DHCP6 is true.
	Nameservers []string `json:"nameservers,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// NTP describes a list of FQDNs, IP4, or IP6 addresses used as NTP servers.
	NTP []string `json:"ntp,omitempty"`

	// +optional

	// PortID identifies the allocated port on the backing network, is only
	// valid on the requested node, and only set if port allocation was
	// requested.
	PortID string `json:"portID,omitempty"`

	// +optional

	// ProviderRef describes the Kubernetes, network-like resource that provides
	// this network interface's underlying capabilities.
	ProviderRef *corev1.TypedLocalObjectReference `json:"providerRef,omitempty"`

	// +optional
	// +kubebuilder:validation:UniqueItems=true

	// SearchDomains describes an optional list of DNS search domains configured
	// for this interface.
	//
	// This value will be empty when DHCP4 and DHCP6 are both true.
	SearchDomains []string `json:"searchDomains,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=vmnic
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Network",type="string",JSONPath=".spec.networkName"
// +kubebuilder:printcolumn:name="IP-Addrs",type="string",JSONPath=".status.addresses[*]"
// +kubebuilder:printcolumn:name="Att-Name",type="string",JSONPath=".spec.attachedTo.name"
// +kubebuilder:printcolumn:name="Att-Kind",type="string",JSONPath=".spec.attachedTo.kind"

// NetworkInterface is the schema for the networkinterfaces API.
type NetworkInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NetworkInterfaceSpec   `json:"spec,omitempty"`
	Status NetworkInterfaceStatus `json:"status,omitempty"`
}

// NamespacedName returns this resource's namespace and name joined with a
// "/" character.
func (r NetworkInterface) NamespacedName() string {
	return r.Namespace + "/" + r.Name
}

// +kubebuilder:object:root=true

// NetworkInterfaceList contains a list of
// NetworkInterface.
type NetworkInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NetworkInterface `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&NetworkInterface{}, &NetworkInterfaceList{})
}
