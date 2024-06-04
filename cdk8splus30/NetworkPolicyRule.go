package cdk8splus30


// Describes a rule allowing traffic from / to pods matched by a network policy selector.
type NetworkPolicyRule struct {
	// Peer this rule interacts with.
	Peer INetworkPolicyPeer `field:"required" json:"peer" yaml:"peer"`
	// The ports of the rule.
	// Default: - traffic is allowed on all ports.
	//
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

