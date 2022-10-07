//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func validateNamespaces_AllParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateNamespaces_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNamespaces_SelectParameters(scope constructs.Construct, id *string, options *NamespacesSelectOptions) error {
	return nil
}

func validateNewNamespacesParameters(scope constructs.Construct, id *string) error {
	return nil
}

