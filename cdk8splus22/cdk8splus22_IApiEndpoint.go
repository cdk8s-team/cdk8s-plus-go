// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// An API Endpoint can either be a resource descriptor (e.g /pods) or a non resource url (e.g /healthz). It must be one or the other, and not both.
type IApiEndpoint interface {
	// Return the IApiResource this object represents.
	AsApiResource() IApiResource
	// Return the non resource url this object represents.
	AsNonApiResource() *string
}

// The jsii proxy for IApiEndpoint
type jsiiProxy_IApiEndpoint struct {
	_ byte // padding
}

func (i *jsiiProxy_IApiEndpoint) AsApiResource() IApiResource {
	var returns IApiResource

	_jsii_.Invoke(
		i,
		"asApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IApiEndpoint) AsNonApiResource() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"asNonApiResource",
		nil, // no parameters
		&returns,
	)

	return returns
}

