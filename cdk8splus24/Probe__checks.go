//go:build !no_runtime_type_checking

package cdk8splus24

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateProbe_FromCommandParameters(command *[]*string, options *CommandProbeOptions) error {
	if command == nil {
		return fmt.Errorf("parameter command is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateProbe_FromHttpGetParameters(path *string, options *HttpGetProbeOptions) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateProbe_FromTcpSocketParameters(options *TcpSocketProbeOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

