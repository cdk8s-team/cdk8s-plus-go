//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodConnections) validateAllowFromParameters(peer INetworkPolicyPeer, options *PodConnectionsAllowFromOptions) error {
	return nil
}

func (p *jsiiProxy_PodConnections) validateAllowToParameters(peer INetworkPolicyPeer, options *PodConnectionsAllowToOptions) error {
	return nil
}

func validateNewPodConnectionsParameters(instance AbstractPod) error {
	return nil
}

