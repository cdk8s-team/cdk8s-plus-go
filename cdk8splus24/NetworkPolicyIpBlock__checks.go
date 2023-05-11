//go:build !no_runtime_type_checking

package cdk8splus24

import (
	"fmt"

	"github.com/aws/constructs-go/constructs/v10"
)

func validateNetworkPolicyIpBlock_AnyIpv4Parameters(scope constructs.Construct, id *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func validateNetworkPolicyIpBlock_AnyIpv6Parameters(scope constructs.Construct, id *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	return nil
}

func validateNetworkPolicyIpBlock_Ipv4Parameters(scope constructs.Construct, id *string, cidrIp *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if cidrIp == nil {
		return fmt.Errorf("parameter cidrIp is required, but nil was provided")
	}

	return nil
}

func validateNetworkPolicyIpBlock_Ipv6Parameters(scope constructs.Construct, id *string, cidrIp *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if cidrIp == nil {
		return fmt.Errorf("parameter cidrIp is required, but nil was provided")
	}

	return nil
}

func validateNetworkPolicyIpBlock_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

