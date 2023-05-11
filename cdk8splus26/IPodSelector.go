package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/internal"
)

// Represents an object that can select pods.
type IPodSelector interface {
	constructs.IConstruct
	// Return the configuration of this selector.
	ToPodSelectorConfig() *PodSelectorConfig
}

// The jsii proxy for IPodSelector
type jsiiProxy_IPodSelector struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_IPodSelector) ToPodSelectorConfig() *PodSelectorConfig {
	var returns *PodSelectorConfig

	_jsii_.Invoke(
		i,
		"toPodSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

