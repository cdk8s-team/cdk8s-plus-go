//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func validateReplicas_AbsoluteParameters(value *float64) error {
	return nil
}

func validateReplicas_PercentParameters(value *float64) error {
	return nil
}

