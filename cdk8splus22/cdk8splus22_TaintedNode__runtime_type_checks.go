//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	"fmt"
)

func validateNewTaintedNodeParameters(taintSelector *[]NodeTaintQuery) error {
	if taintSelector == nil {
		return fmt.Errorf("parameter taintSelector is required, but nil was provided")
	}

	return nil
}

