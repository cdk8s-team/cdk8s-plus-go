// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/internal"
)

// Represents an object that can be used as a role binding subject.
type ISubject interface {
	constructs.IConstruct
	// Return the subject configuration.
	ToSubjectConfiguration() *SubjectConfiguration
}

// The jsii proxy for ISubject
type jsiiProxy_ISubject struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_ISubject) ToSubjectConfiguration() *SubjectConfiguration {
	var returns *SubjectConfiguration

	_jsii_.Invoke(
		i,
		"toSubjectConfiguration",
		nil, // no parameters
		&returns,
	)

	return returns
}

