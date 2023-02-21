//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterRoleBinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewClusterRoleBindingParameters(scope constructs.Construct, id *string, props *ClusterRoleBindingProps) error {
	return nil
}

