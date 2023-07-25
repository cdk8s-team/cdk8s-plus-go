package cdk8splus27


// Describes a rule allowing traffic from / to pods matched by a network policy selector.
type NetworkPolicyRule struct {
	// Peer this rule interacts with.
	Peer INetworkPolicyPeer `field:"required" json:"peer" yaml:"peer"`
	// The ports of the rule.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

