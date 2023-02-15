//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

// Building without runtime type checking enabled, so all the below just return nil

func validateNodeTaintQuery_ExistsParameters(key *string, options *NodeTaintQueryOptions) error {
	return nil
}

func validateNodeTaintQuery_IsParameters(key *string, value *string, options *NodeTaintQueryOptions) error {
	return nil
}

