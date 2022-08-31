//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

// Building without runtime type checking enabled, so all the below just return nil

func validateResource_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewResourceParameters(scope constructs.Construct, id *string) error {
	return nil
}

