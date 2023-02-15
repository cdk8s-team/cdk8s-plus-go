//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

// Building without runtime type checking enabled, so all the below just return nil

func validateNetworkPolicyPort_OfParameters(props *NetworkPolicyPortProps) error {
	return nil
}

func validateNetworkPolicyPort_TcpParameters(port *float64) error {
	return nil
}

func validateNetworkPolicyPort_TcpRangeParameters(startPort *float64, endPort *float64) error {
	return nil
}

func validateNetworkPolicyPort_UdpParameters(port *float64) error {
	return nil
}

func validateNetworkPolicyPort_UdpRangeParameters(startPort *float64, endPort *float64) error {
	return nil
}

