package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/internal"
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

