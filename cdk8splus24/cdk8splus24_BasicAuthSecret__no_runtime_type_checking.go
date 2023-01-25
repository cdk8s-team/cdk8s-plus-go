//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func (b *jsiiProxy_BasicAuthSecret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (b *jsiiProxy_BasicAuthSecret) validateEnvValueParameters(key *string, options *EnvValueFromSecretOptions) error {
	return nil
}

func (b *jsiiProxy_BasicAuthSecret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateBasicAuthSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateBasicAuthSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewBasicAuthSecretParameters(scope constructs.Construct, id *string, props *BasicAuthSecretProps) error {
	return nil
}

