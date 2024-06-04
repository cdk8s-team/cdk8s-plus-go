package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

// Represents a scalable workload.
type IScalable interface {
	// Called on all IScalable targets when they are associated with an autoscaler.
	MarkHasAutoscaler()
	// Return the target spec properties of this Scalable.
	ToScalingTarget() *ScalingTarget
	// If this is a target of an autoscaler.
	HasAutoscaler() *bool
	SetHasAutoscaler(h *bool)
}

// The jsii proxy for IScalable
type jsiiProxy_IScalable struct {
	_ byte // padding
}

func (i *jsiiProxy_IScalable) MarkHasAutoscaler() {
	_jsii_.InvokeVoid(
		i,
		"markHasAutoscaler",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IScalable) ToScalingTarget() *ScalingTarget {
	var returns *ScalingTarget

	_jsii_.Invoke(
		i,
		"toScalingTarget",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_IScalable) HasAutoscaler() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"hasAutoscaler",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IScalable)SetHasAutoscaler(val *bool) {
	if err := j.validateSetHasAutoscalerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hasAutoscaler",
		val,
	)
}

