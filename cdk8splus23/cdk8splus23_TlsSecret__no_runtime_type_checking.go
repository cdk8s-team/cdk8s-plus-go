//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TlsSecret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (t *jsiiProxy_TlsSecret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateTlsSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateTlsSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTlsSecretParameters(scope constructs.Construct, id *string, props *TlsSecretProps) error {
	return nil
}

