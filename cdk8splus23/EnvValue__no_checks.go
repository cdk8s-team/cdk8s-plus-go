//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

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

