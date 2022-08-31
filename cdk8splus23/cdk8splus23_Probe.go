// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
)

// Probe describes a health check to be performed against a container to determine whether it is alive or ready to receive traffic.
type Probe interface {
}

// The jsii proxy struct for Probe
type jsiiProxy_Probe struct {
	_ byte // padding
}

// Defines a probe based on a command which is executed within the container.
func Probe_FromCommand(command *[]*string, options *CommandProbeOptions) Probe {
	_init_.Initialize()

	if err := validateProbe_FromCommandParameters(command, options); err != nil {
		panic(err)
	}
	var returns Probe

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Probe",
		"fromCommand",
		[]interface{}{command, options},
		&returns,
	)

	return returns
}

// Defines a probe based on an HTTP GET request to the IP address of the container.
func Probe_FromHttpGet(path *string, options *HttpGetProbeOptions) Probe {
	_init_.Initialize()

	if err := validateProbe_FromHttpGetParameters(path, options); err != nil {
		panic(err)
	}
	var returns Probe

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Probe",
		"fromHttpGet",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Defines a probe based opening a connection to a TCP socket on the container.
func Probe_FromTcpSocket(options *TcpSocketProbeOptions) Probe {
	_init_.Initialize()

	if err := validateProbe_FromTcpSocketParameters(options); err != nil {
		panic(err)
	}
	var returns Probe

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Probe",
		"fromTcpSocket",
		[]interface{}{options},
		&returns,
	)

	return returns
}

