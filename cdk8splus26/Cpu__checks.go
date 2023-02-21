//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

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

