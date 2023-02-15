// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus25/v2/jsii"
)

// Deployment strategies.
type DeploymentStrategy interface {
}

// The jsii proxy struct for DeploymentStrategy
type jsiiProxy_DeploymentStrategy struct {
	_ byte // padding
}

// All existing Pods are killed before new ones are created.
// See: https://kubernetes.io/docs/concepts/workloads/controllers/deployment/#recreate-deployment
//
func DeploymentStrategy_Recreate() DeploymentStrategy {
	_init_.Initialize()

	var returns DeploymentStrategy

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.DeploymentStrategy",
		"recreate",
		nil, // no parameters
		&returns,
	)

	return returns
}

func DeploymentStrategy_RollingUpdate(options *DeploymentStrategyRollingUpdateOptions) DeploymentStrategy {
	_init_.Initialize()

	if err := validateDeploymentStrategy_RollingUpdateParameters(options); err != nil {
		panic(err)
	}
	var returns DeploymentStrategy

	_jsii_.StaticInvoke(
		"cdk8s-plus-25.DeploymentStrategy",
		"rollingUpdate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

