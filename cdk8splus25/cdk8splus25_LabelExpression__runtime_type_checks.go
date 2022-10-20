//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	"fmt"
)

func validateLabelExpression_DoesNotExistParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateLabelExpression_ExistsParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateLabelExpression_InParameters(key *string, values *[]*string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateLabelExpression_NotInParameters(key *string, values *[]*string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

