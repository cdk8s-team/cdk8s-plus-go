//go:build !no_runtime_type_checking

package cdk8splus29

import (
	"fmt"
)

func validateMetricTarget_AverageUtilizationParameters(averageUtilization *float64) error {
	if averageUtilization == nil {
		return fmt.Errorf("parameter averageUtilization is required, but nil was provided")
	}

	return nil
}

func validateMetricTarget_AverageValueParameters(averageValue *float64) error {
	if averageValue == nil {
		return fmt.Errorf("parameter averageValue is required, but nil was provided")
	}

	return nil
}

func validateMetricTarget_ValueParameters(value *float64) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

