//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceAccount) validateAddSecretParameters(secr ISecret) error {
	return nil
}

func validateServiceAccount_FromServiceAccountNameParameters(scope constructs.Construct, id *string, name *string, options *FromServiceAccountNameOptions) error {
	return nil
}

func validateServiceAccount_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewServiceAccountParameters(scope constructs.Construct, id *string, props *ServiceAccountProps) error {
	return nil
}

