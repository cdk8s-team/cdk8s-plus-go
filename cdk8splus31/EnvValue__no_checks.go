//go:build no_runtime_type_checking

package cdk8splus31

// Building without runtime type checking enabled, so all the below just return nil

func validateEnvValue_FromConfigMapParameters(configMap IConfigMap, key *string, options *EnvValueFromConfigMapOptions) error {
	return nil
}

func validateEnvValue_FromFieldRefParameters(fieldPath EnvFieldPaths, options *EnvValueFromFieldRefOptions) error {
	return nil
}

func validateEnvValue_FromProcessParameters(key *string, options *EnvValueFromProcessOptions) error {
	return nil
}

func validateEnvValue_FromResourceParameters(resource ResourceFieldPaths, options *EnvValueFromResourceOptions) error {
	return nil
}

func validateEnvValue_FromSecretValueParameters(secretValue *SecretValue, options *EnvValueFromSecretOptions) error {
	return nil
}

func validateEnvValue_FromValueParameters(value *string) error {
	return nil
}

