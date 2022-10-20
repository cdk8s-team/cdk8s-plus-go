//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

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

