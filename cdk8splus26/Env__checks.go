//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	"fmt"
)

func (e *jsiiProxy_Env) validateAddVariableParameters(name *string, value EnvValue) error {
	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (e *jsiiProxy_Env) validateCopyFromParameters(from EnvFrom) error {
	if from == nil {
		return fmt.Errorf("parameter from is required, but nil was provided")
	}

	return nil
}

func validateEnv_FromConfigMapParameters(configMap IConfigMap) error {
	if configMap == nil {
		return fmt.Errorf("parameter configMap is required, but nil was provided")
	}

	return nil
}

func validateEnv_FromSecretParameters(secr ISecret) error {
	if secr == nil {
		return fmt.Errorf("parameter secr is required, but nil was provided")
	}

	return nil
}

func validateNewEnvParameters(sources *[]EnvFrom, variables *map[string]EnvValue) error {
	if sources == nil {
		return fmt.Errorf("parameter sources is required, but nil was provided")
	}

	if variables == nil {
		return fmt.Errorf("parameter variables is required, but nil was provided")
	}

	return nil
}

