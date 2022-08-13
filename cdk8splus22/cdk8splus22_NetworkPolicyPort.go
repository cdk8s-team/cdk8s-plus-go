// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"
)

// Describes a port to allow traffic on.
type NetworkPolicyPort interface {
}

// The jsii proxy struct for NetworkPolicyPort
type jsiiProxy_NetworkPolicyPort struct {
	_ byte // padding
}

// Any TCP traffic.
func NetworkPolicyPort_AllTcp() NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NetworkPolicyPort",
		"allTcp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Any UDP traffic.
func NetworkPolicyPort_AllUdp() NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NetworkPolicyPort",
		"allUdp",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Custom port configuration.
func NetworkPolicyPort_Of(props *NetworkPolicyPortProps) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NetworkPolicyPort",
		"of",
		[]interface{}{props},
		&returns,
	)

	return returns
}

// Distinct TCP ports.
func NetworkPolicyPort_Tcp(port *float64) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NetworkPolicyPort",
		"tcp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// A TCP port range.
func NetworkPolicyPort_TcpRange(startPort *float64, endPort *float64) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NetworkPolicyPort",
		"tcpRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

// Distinct UDP ports.
func NetworkPolicyPort_Udp(port *float64) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NetworkPolicyPort",
		"udp",
		[]interface{}{port},
		&returns,
	)

	return returns
}

// A UDP port range.
func NetworkPolicyPort_UdpRange(startPort *float64, endPort *float64) NetworkPolicyPort {
	_init_.Initialize()

	var returns NetworkPolicyPort

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NetworkPolicyPort",
		"udpRange",
		[]interface{}{startPort, endPort},
		&returns,
	)

	return returns
}

