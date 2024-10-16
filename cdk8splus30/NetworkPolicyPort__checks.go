//go:build !no_runtime_type_checking

package cdk8splus30

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateNetworkPolicyPort_OfParameters(props *NetworkPolicyPortProps) error {
	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

func validateNetworkPolicyPort_TcpParameters(port *float64) error {
	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

func validateNetworkPolicyPort_TcpRangeParameters(startPort *float64, endPort *float64) error {
	if startPort == nil {
		return fmt.Errorf("parameter startPort is required, but nil was provided")
	}

	if endPort == nil {
		return fmt.Errorf("parameter endPort is required, but nil was provided")
	}

	return nil
}

func validateNetworkPolicyPort_UdpParameters(port *float64) error {
	if port == nil {
		return fmt.Errorf("parameter port is required, but nil was provided")
	}

	return nil
}

func validateNetworkPolicyPort_UdpRangeParameters(startPort *float64, endPort *float64) error {
	if startPort == nil {
		return fmt.Errorf("parameter startPort is required, but nil was provided")
	}

	if endPort == nil {
		return fmt.Errorf("parameter endPort is required, but nil was provided")
	}

	return nil
}

