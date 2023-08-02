package cdk8splus27


// Describes how the network policy should configure egress / ingress traffic.
type NetworkPolicyTraffic struct {
	// Specifies the default behavior of the policy when no rules are defined.
	// Default: - unset, the policy does not change the behavior.
	//
	Default NetworkPolicyTrafficDefault `field:"optional" json:"default" yaml:"default"`
	// List of rules to be applied to the selected pods.
	//
	// If empty, the behavior of the policy is dictated by the `default` property.
	// Default: - no rules.
	//
	Rules *[]*NetworkPolicyRule `field:"optional" json:"rules" yaml:"rules"`
}

