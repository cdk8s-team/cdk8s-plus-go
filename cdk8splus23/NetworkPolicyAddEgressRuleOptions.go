// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Options for `NetworkPolicy.addEgressRule`.
type NetworkPolicyAddEgressRuleOptions struct {
	// Ports the rule should allow outgoing traffic to.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

