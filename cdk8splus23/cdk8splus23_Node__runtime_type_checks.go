//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	"fmt"
)

func validateNode_NamedParameters(nodeName *string) error {
	if nodeName == nil {
		return fmt.Errorf("parameter nodeName is required, but nil was provided")
	}

	return nil
}

