//go:build no_runtime_type_checking

package cdk8splus26

// Building without runtime type checking enabled, so all the below just return nil

func validateCpu_MillisParameters(amount *float64) error {
	return nil
}

func validateCpu_UnitsParameters(amount *float64) error {
	return nil
}

func (j *jsiiProxy_Cpu) validateSetAmountParameters(val *string) error {
	return nil
}

