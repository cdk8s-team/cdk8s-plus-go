//go:build no_runtime_type_checking

package cdk8splus27

// Building without runtime type checking enabled, so all the below just return nil

func validateMetricTarget_AverageUtilizationParameters(averageUtilization *float64) error {
	return nil
}

func validateMetricTarget_AverageValueParameters(averageValue *float64) error {
	return nil
}

func validateMetricTarget_ValueParameters(value *float64) error {
	return nil
}

