//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func validateNetworkPolicyIpBlock_AnyIpv4Parameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateNetworkPolicyIpBlock_AnyIpv6Parameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateNetworkPolicyIpBlock_Ipv4Parameters(scope constructs.Construct, id *string, cidrIp *string) error {
	return nil
}

func validateNetworkPolicyIpBlock_Ipv6Parameters(scope constructs.Construct, id *string, cidrIp *string) error {
	return nil
}

func validateNetworkPolicyIpBlock_IsConstructParameters(x interface{}) error {
	return nil
}

