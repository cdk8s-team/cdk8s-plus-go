// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Describes a rule allowing traffic from / to pods matched by a network policy selector.
type NetworkPolicyRule struct {
	// Peer this rule interacts with.
	Peer INetworkPolicyPeer `field:"required" json:"peer" yaml:"peer"`
	// The ports of the rule.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

