//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateDeploymentStrategy_RollingUpdateParameters(options *DeploymentStrategyRollingUpdateOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

