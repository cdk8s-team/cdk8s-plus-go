//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Secret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_Secret) validateEnvValueParameters(key *string, options *EnvValueFromSecretOptions) error {
	return nil
}

func (s *jsiiProxy_Secret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSecretParameters(scope constructs.Construct, id *string, props *SecretProps) error {
	return nil
}

