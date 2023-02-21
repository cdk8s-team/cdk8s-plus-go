//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_ClusterRole) validateAggregateParameters(key *string, value *string) error {
	if key == nil {
		return fmt.Errorf("parameter key is required, but nil was provided")
	}

	if value == nil {
		return fmt.Errorf("parameter value is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ClusterRole) validateAllowParameters(verbs *[]*string) error {
	if verbs == nil {
		return fmt.Errorf("parameter verbs is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ClusterRole) validateBindInNamespaceParameters(namespace *string) error {
	if namespace == nil {
		return fmt.Errorf("parameter namespace is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_ClusterRole) validateCombineParameters(rol ClusterRole) error {
	if rol == nil {
		return fmt.Errorf("parameter rol is required, but nil was provided")
	}

	return nil
}

func validateClusterRole_FromClusterRoleNameParameters(scope constructs.Construct, id *string, name *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if name == nil {
		return fmt.Errorf("parameter name is required, but nil was provided")
	}

	return nil
}

func validateClusterRole_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewClusterRoleParameters(scope constructs.Construct, id *string, props *ClusterRoleProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

