//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateEnvValue_FromConfigMapParameters(configMap IConfigMap, key *string, options *EnvValueFromConfigMapOptions) error {
	if configMap == nil {
		return fmt.Errorf("parameter configMap is required, but nil was provided")
	}

	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEnvValue_FromFieldRefParameters(fieldPath EnvFieldPaths, options *EnvValueFromFieldRefOptions) error {
	if fieldPath == "" {
		return fmt.Errorf("parameter fieldPath is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEnvValue_FromProcessParameters(key *string, options *EnvValueFromProcessOptions) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEnvValue_FromResourceParameters(resource ResourceFieldPaths, options *EnvValueFromResourceOptions) error {
	if resource == "" {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEnvValue_FromSecretValueParameters(secretValue *SecretValue, options *EnvValueFromSecretOptions) error {
	if secretValue == nil {
		return fmt.Errorf("parameter secretValue is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(secretValue, func() string { return "parameter secretValue" }); err != nil {
		return err
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateEnvValue_FromValueParameters(value *string) error {
	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

