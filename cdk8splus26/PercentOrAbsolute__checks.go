//go:build !no_runtime_type_checking

package cdk8splus26

import (
	"fmt"
)

func validatePercentOrAbsolute_AbsoluteParameters(num *float64) error {
	if num == nil {
		return fmt.Errorf("parameter num is required, but nil was provided")
	}

	return nil
}

func validatePercentOrAbsolute_PercentParameters(percent *float64) error {
	if percent == nil {
		return fmt.Errorf("parameter percent is required, but nil was provided")
	}

	return nil
}

