// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Default behaviors of network traffic in policies.
type NetworkPolicyTrafficDefault string

const (
	// The policy denies all traffic.
	//
	// Since rules are additive, additional rules or policies can allow
	// specific traffic.
	NetworkPolicyTrafficDefault_DENY NetworkPolicyTrafficDefault = "DENY"
	// The policy allows all traffic (either ingress or egress).
	//
	// Since rules are additive, no additional rule or policies can
	// subsequently deny the traffic.
	NetworkPolicyTrafficDefault_ALLOW NetworkPolicyTrafficDefault = "ALLOW"
)

