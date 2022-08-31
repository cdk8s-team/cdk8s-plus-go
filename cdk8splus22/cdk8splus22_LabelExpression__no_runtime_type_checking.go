//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

// Building without runtime type checking enabled, so all the below just return nil

func validateLabelExpression_DoesNotExistParameters(key *string) error {
	return nil
}

func validateLabelExpression_ExistsParameters(key *string) error {
	return nil
}

func validateLabelExpression_InParameters(key *string, values *[]*string) error {
	return nil
}

func validateLabelExpression_NotInParameters(key *string, values *[]*string) error {
	return nil
}

