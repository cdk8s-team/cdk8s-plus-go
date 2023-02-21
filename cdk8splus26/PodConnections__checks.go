//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (p *jsiiProxy_PodConnections) validateAllowFromParameters(peer INetworkPolicyPeer, options *PodConnectionsAllowFromOptions) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PodConnections) validateAllowToParameters(peer INetworkPolicyPeer, options *PodConnectionsAllowToOptions) error {
	if peer == nil {
		return fmt.Errorf("parameter peer is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func validateNewPodConnectionsParameters(instance AbstractPod) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

