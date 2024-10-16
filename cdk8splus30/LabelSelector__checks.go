//go:build !no_runtime_type_checking

package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateLabelSelector_OfParameters(options *LabelSelectorOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

