//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateHandler_FromCommandParameters(command *[]*string) error {
	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	return nil
}

func validateHandler_FromHttpGetParameters(path *string, options *HandlerFromHttpGetOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateHandler_FromTcpSocketParameters(options *HandlerFromTcpSocketOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

