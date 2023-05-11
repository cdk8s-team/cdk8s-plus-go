//go:build no_runtime_type_checking

package cdk8splus26

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

