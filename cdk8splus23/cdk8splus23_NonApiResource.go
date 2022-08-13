// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
)

// Factory for creating non api resources.
type NonApiResource interface {
	IApiEndpoint
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
}

// The jsii proxy struct for NonApiResource
type jsiiProxy_NonApiResource struct {
	jsiiProxy_IApiEndpoint
}

func NonApiResource_Of(url *string) NonApiResource {
	_init_.Initialize()

	var returns NonApiResource

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.NonApiResource",
		"of",
		[]interface{}{url},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NonApiResource) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		n,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NonApiResource) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

