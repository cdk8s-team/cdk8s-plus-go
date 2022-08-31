//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

// Building without runtime type checking enabled, so all the below just return nil

func validateNodeLabelQuery_DoesNotExistParameters(key *string) error {
	return nil
}

func validateNodeLabelQuery_ExistsParameters(key *string) error {
	return nil
}

func validateNodeLabelQuery_GtParameters(key *string, values *[]*string) error {
	return nil
}

func validateNodeLabelQuery_InParameters(key *string, values *[]*string) error {
	return nil
}

func validateNodeLabelQuery_IsParameters(key *string, value *string) error {
	return nil
}

func validateNodeLabelQuery_LtParameters(key *string, values *[]*string) error {
	return nil
}

func validateNodeLabelQuery_NotInParameters(key *string, values *[]*string) error {
	return nil
}

