//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (c *jsiiProxy_Container) validateAddPortParameters(port *ContainerPort) error {
	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(port, func() string { return "parameter port" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_Container) validateMountParameters(path *string, storage IStorage, options *MountOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if storage == nil {
		return fmt.Errorf("parameter storage is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewContainerParameters(props *ContainerProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

