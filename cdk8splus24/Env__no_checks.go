//go:build no_runtime_type_checking

package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func (e *jsiiProxy_Env) validateAddVariableParameters(name *string, value EnvValue) error {
	return nil
}

func (e *jsiiProxy_Env) validateCopyFromParameters(from EnvFrom) error {
	return nil
}

func validateEnv_FromConfigMapParameters(configMap IConfigMap) error {
	return nil
}

func validateEnv_FromSecretParameters(secr ISecret) error {
	return nil
}

func validateNewEnvParameters(sources *[]EnvFrom, variables *map[string]EnvValue) error {
	return nil
}

