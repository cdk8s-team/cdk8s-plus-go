// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Configuration for network peers.
//
// A peer can either by an ip block, or a selection of pods, not both.
type NetworkPolicyPeerConfig struct {
	// The ip block this peer represents.
	IpBlock NetworkPolicyIpBlock `field:"optional" json:"ipBlock" yaml:"ipBlock"`
	// The pod selector this peer represents.
	PodSelector *PodSelectorConfig `field:"optional" json:"podSelector" yaml:"podSelector"`
}

