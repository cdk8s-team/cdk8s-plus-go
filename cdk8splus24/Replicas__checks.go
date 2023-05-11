//go:build !no_runtime_type_checking

package cdk8splus24

import (
	"fmt"
)

func validateReplicas_AbsoluteParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateReplicas_PercentParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

