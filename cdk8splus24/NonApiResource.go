package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"
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

	if err := validateNonApiResource_OfParameters(url); err != nil {
		panic(err)
	}
	var returns NonApiResource

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NonApiResource",
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

