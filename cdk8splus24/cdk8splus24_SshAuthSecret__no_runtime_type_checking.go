//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SshAuthSecret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_SshAuthSecret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateSshAuthSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateSshAuthSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSshAuthSecretParameters(scope constructs.Construct, id *string, props *SshAuthSecretProps) error {
	return nil
}

