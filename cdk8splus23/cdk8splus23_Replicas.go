// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
)

// The amount of replicas that will change.
type Replicas interface {
}

// The jsii proxy struct for Replicas
type jsiiProxy_Replicas struct {
	_ byte // padding
}

// Changes the pods by a percentage of the it's current value.
func Replicas_Absolute(value *float64) Replicas {
	_init_.Initialize()

	if err := validateReplicas_AbsoluteParameters(value); err != nil {
		panic(err)
	}
	var returns Replicas

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Replicas",
		"absolute",
		[]interface{}{value},
		&returns,
	)

	return returns
}

// Changes the pods by a percentage of the it's current value.
func Replicas_Percent(value *float64) Replicas {
	_init_.Initialize()

	if err := validateReplicas_PercentParameters(value); err != nil {
		panic(err)
	}
	var returns Replicas

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Replicas",
		"percent",
		[]interface{}{value},
		&returns,
	)

	return returns
}

