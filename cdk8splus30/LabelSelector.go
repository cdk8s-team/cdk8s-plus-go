package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"
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

	if err := validateLabelSelector_OfParameters(options); err != nil {
		panic(err)
	}
	var returns LabelSelector

	_jsii_.StaticInvoke(
		"cdk8s-plus-30.LabelSelector",
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

