//go:build !no_runtime_type_checking

package cdk8splus25

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateStatefulSetUpdateStrategy_RollingUpdateParameters(options *StatefulSetUpdateStrategyRollingUpdateOptions) error {
	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

