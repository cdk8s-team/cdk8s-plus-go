// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
)

// StatefulSet update strategies.
type StatefulSetUpdateStrategy interface {
}

// The jsii proxy struct for StatefulSetUpdateStrategy
type jsiiProxy_StatefulSetUpdateStrategy struct {
	_ byte // padding
}

// The controller will not automatically update the Pods in a StatefulSet.
//
// Users must manually delete Pods to cause the controller to create new Pods
// that reflect modifications.
func StatefulSetUpdateStrategy_OnDelete() StatefulSetUpdateStrategy {
	_init_.Initialize()

	var returns StatefulSetUpdateStrategy

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.StatefulSetUpdateStrategy",
		"onDelete",
		nil, // no parameters
		&returns,
	)

	return returns
}

// The controller will delete and recreate each Pod in the StatefulSet.
//
// It will proceed in the same order as Pod termination (from the largest ordinal to the smallest),
// updating each Pod one at a time. The Kubernetes control plane waits until an updated
// Pod is Running and Ready prior to updating its predecessor.
func StatefulSetUpdateStrategy_RollingUpdate(options *StatefulSetUpdateStrategyRollingUpdateOptions) StatefulSetUpdateStrategy {
	_init_.Initialize()

	var returns StatefulSetUpdateStrategy

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.StatefulSetUpdateStrategy",
		"rollingUpdate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

