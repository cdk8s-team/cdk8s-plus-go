//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

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

