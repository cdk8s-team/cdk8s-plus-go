//go:build no_runtime_type_checking

package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func validatePercentOrAbsolute_AbsoluteParameters(num *float64) error {
	return nil
}

func validatePercentOrAbsolute_PercentParameters(percent *float64) error {
	return nil
}

