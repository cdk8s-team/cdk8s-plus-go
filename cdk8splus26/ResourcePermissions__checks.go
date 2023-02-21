//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	"fmt"
)

func validateNewResourcePermissionsParameters(instance Resource) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

