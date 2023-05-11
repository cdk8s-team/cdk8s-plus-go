package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"
)

// Defines a specific action that should be taken.
type Handler interface {
}

// The jsii proxy struct for Handler
type jsiiProxy_Handler struct {
	_ byte // padding
}

// Defines a handler based on a command which is executed within the container.
func Handler_FromCommand(command *[]*string) Handler {
	_init_.Initialize()

	if err := validateHandler_FromCommandParameters(command); err != nil {
		panic(err)
	}
	var returns Handler

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Handler",
		"fromCommand",
		[]interface{}{command},
		&returns,
	)

	return returns
}

// Defines a handler based on an HTTP GET request to the IP address of the container.
func Handler_FromHttpGet(path *string, options *HandlerFromHttpGetOptions) Handler {
	_init_.Initialize()

	if err := validateHandler_FromHttpGetParameters(path, options); err != nil {
		panic(err)
	}
	var returns Handler

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Handler",
		"fromHttpGet",
		[]interface{}{path, options},
		&returns,
	)

	return returns
}

// Defines a handler based opening a connection to a TCP socket on the container.
func Handler_FromTcpSocket(options *HandlerFromTcpSocketOptions) Handler {
	_init_.Initialize()

	if err := validateHandler_FromTcpSocketParameters(options); err != nil {
		panic(err)
	}
	var returns Handler

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Handler",
		"fromTcpSocket",
		[]interface{}{options},
		&returns,
	)

	return returns
}

