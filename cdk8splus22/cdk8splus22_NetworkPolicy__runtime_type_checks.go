//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (n *jsiiProxy_NetworkPolicy) validateAddEgressRuleParameters(peer INetworkPolicyPeer) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateAddIngressRuleParameters(peer INetworkPolicyPeer) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	return nil
}

func validateNetworkPolicy_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewNetworkPolicyParameters(scope constructs.Construct, id *string, props *NetworkPolicyProps) error {
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

