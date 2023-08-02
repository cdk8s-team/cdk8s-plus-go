package cdk8splus27


// Options for `NetworkPolicy.addEgressRule`.
type NetworkPolicyAddEgressRuleOptions struct {
	// Ports the rule should allow outgoing traffic to.
	// Default: - If the peer is a managed pod, take its ports. Otherwise, all ports are allowed.
	//
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

