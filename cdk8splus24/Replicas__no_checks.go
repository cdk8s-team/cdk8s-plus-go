//go:build no_runtime_type_checking

package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func validateReplicas_AbsoluteParameters(value *float64) error {
	return nil
}

func validateReplicas_PercentParameters(value *float64) error {
	return nil
}

