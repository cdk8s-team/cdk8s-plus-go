package cdk8splus25


// Options for `NetworkPolicy.addEgressRule`.
type NetworkPolicyAddEgressRuleOptions struct {
	// Ports the rule should allow outgoing traffic to.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

