package cdk8splus28


// Configuration for network peers.
//
// A peer can either by an ip block, or a selection of pods, not both.
type NetworkPolicyPeerConfig struct {
	// The ip block this peer represents.
	IpBlock NetworkPolicyIpBlock `field:"optional" json:"ipBlock" yaml:"ipBlock"`
	// The pod selector this peer represents.
	PodSelector *PodSelectorConfig `field:"optional" json:"podSelector" yaml:"podSelector"`
}

