//go:build !no_runtime_type_checking

package cdk8splus29

import (
	"fmt"
)

func validateCpu_MillisParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func validateCpu_UnitsParameters(amount *float64) error {
	if amount == nil {
		return fmt.Errorf("parameter amount is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_Cpu) validateSetAmountParameters(val *string) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

