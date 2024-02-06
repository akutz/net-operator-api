// Copyright (c) 2024 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// IPAddressSpec defines the desired state of a IPAddress resource.
type IPAddressSpec struct {

	// +optional

	// Address describes the desired IP address.
	//
	// If specified, the IP address must be within the range of the referenced
	// Network.
	//
	// If omitted, an IP address is allocated from the configured IP range of
	// the referenced Network.
	Address string `json:"address,omitempty"`

	// +optional
	// +kubebuilder:default:=ip4

	// Family describes the IP family of the requested address.
	Family IPFamily `json:"family,omitempty"`

	// +optional

	// NetworkName describes the name of the Network resource from which to
	// allocate an IP address.
	//
	// If omitted, this field defaults to the name of the namespace's default
	// network. The default network for a namespace is either the sole Network
	// resource in the namespace, or if multiple Network resources, the one with
	// the label netoperator.vmware.com/default-network.
	NetworkName string `json:"networkName,omitempty"`

	// +optional

	// ProviderConfig describes additional configuration information for the
	// underlying network provider.
	ProviderConfig *ProviderConfig `json:"providerConfig,omitempty"`
}

// IPAddressStatus defines the observed state of a IPAddress resource.
type IPAddressStatus struct {

	// +optional

	// Address describes the realized IP address.
	Address string `json:"address,omitempty"`

	// +optional

	// NetworkInterfaceName describes the name of the NetworkInterface using
	// this address.
	//
	// This value is updated whenever the IPAddress is associated with a
	// NetworkInterface. This value is also cleared when an IPAddress is no
	// longer associated with a NetworkInterface.
	NetworkInterfaceName string `json:"networkInterfaceName,omitempty"`

	// +optional

	// Conditions describes the observed conditions of the resource.
	Conditions []metav1.Condition `json:"conditions,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=ip
// +kubebuilder:storageversion
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="Address",type="string",JSONPath=".status.address"
// +kubebuilder:printcolumn:name="Family",type="string",JSONPath=".spec.family"

// Network is the schema for the network API.
type IPAddress struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   IPAddressSpec   `json:"spec,omitempty"`
	Status IPAddressStatus `json:"status,omitempty"`
}

// NamespacedName returns this resource's namespace and name joined with a
// "/" character.
func (r IPAddress) NamespacedName() string {
	return r.Namespace + "/" + r.Name
}

// +kubebuilder:object:root=true

// IPAddressList contains a list of Network.
type IPAddressList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []IPAddress `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&IPAddress{}, &IPAddressList{})
}
