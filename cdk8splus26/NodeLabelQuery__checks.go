//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	"fmt"
)

func validateNodeLabelQuery_DoesNotExistParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateNodeLabelQuery_ExistsParameters(key *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	return nil
}

func validateNodeLabelQuery_GtParameters(key *string, values *[]*string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateNodeLabelQuery_InParameters(key *string, values *[]*string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateNodeLabelQuery_IsParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func validateNodeLabelQuery_LtParameters(key *string, values *[]*string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

func validateNodeLabelQuery_NotInParameters(key *string, values *[]*string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if values == nil {
		return fmt.Errorf("parameter values is required, but nil was provided")
	}

	return nil
}

