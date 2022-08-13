// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Describes how the network policy should configure egress / ingress traffic.
type NetworkPolicyTraffic struct {
	// Specifies the default behavior of the policy when no rules are defined.
	Default NetworkPolicyTrafficDefault `field:"optional" json:"default" yaml:"default"`
	// List of rules to be applied to the selected pods.
	//
	// If empty, the behavior of the policy is dictated by the `default` property.
	Rules *[]*NetworkPolicyRule `field:"optional" json:"rules" yaml:"rules"`
}

