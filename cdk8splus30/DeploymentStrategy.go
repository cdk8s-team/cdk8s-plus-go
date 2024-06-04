package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"
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
		"cdk8s-plus-30.DeploymentStrategy",
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
		"cdk8s-plus-30.DeploymentStrategy",
		"rollingUpdate",
		[]interface{}{options},
		&returns,
	)

	return returns
}

