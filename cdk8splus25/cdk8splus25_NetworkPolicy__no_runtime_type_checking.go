//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func (n *jsiiProxy_NetworkPolicy) validateAddEgressRuleParameters(peer INetworkPolicyPeer) error {
	return nil
}

func (n *jsiiProxy_NetworkPolicy) validateAddIngressRuleParameters(peer INetworkPolicyPeer) error {
	return nil
}

func validateNetworkPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNetworkPolicyParameters(scope constructs.Construct, id *string, props *NetworkPolicyProps) error {
	return nil
}

