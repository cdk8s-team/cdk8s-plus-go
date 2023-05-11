package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a resource or collection of resources.
type IApiResource interface {
	// The group portion of the API version (e.g. `authorization.k8s.io`).
	ApiGroup() *string
	// The unique, namespace-global, name of an object inside the Kubernetes cluster.
	//
	// If this is omitted, the ApiResource should represent all objects of the given type.
	ResourceName() *string
	// The name of a resource type as it appears in the relevant API endpoint.
	//
	// Example:
	//   - "pods" or "pods/log"
	//
	// See: https://kubernetes.io/docs/reference/access-authn-authz/rbac/#referring-to-resources
	//
	ResourceType() *string
}

// The jsii proxy for IApiResource
type jsiiProxy_IApiResource struct {
	_ byte // padding
}

func (j *jsiiProxy_IApiResource) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiResource) ResourceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IApiResource) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

