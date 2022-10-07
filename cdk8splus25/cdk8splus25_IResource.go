// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/internal"
)

// Represents a resource.
type IResource interface {
	constructs.IConstruct
	// The group portion of the API version (e.g. "authorization.k8s.io").
	ApiGroup() *string
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion() *string
	// The object kind (e.g. "Deployment").
	Kind() *string
	// The Kubernetes name of this resource.
	Name() *string
}

// The jsii proxy for IResource
type jsiiProxy_IResource struct {
	internal.Type__constructsIConstruct
}

func (j *jsiiProxy_IResource) ApiGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) ApiVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) Kind() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IResource) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

