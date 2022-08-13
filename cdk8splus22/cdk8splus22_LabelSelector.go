// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"
)

// Match a resource by labels.
type LabelSelector interface {
	IsEmpty() *bool
}

// The jsii proxy struct for LabelSelector
type jsiiProxy_LabelSelector struct {
	_ byte // padding
}

func LabelSelector_Of(options *LabelSelectorOptions) LabelSelector {
	_init_.Initialize()

	var returns LabelSelector

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.LabelSelector",
		"of",
		[]interface{}{options},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabelSelector) IsEmpty() *bool {
	var returns *bool

	_jsii_.Invoke(
		l,
		"isEmpty",
		nil, // no parameters
		&returns,
	)

	return returns
}

