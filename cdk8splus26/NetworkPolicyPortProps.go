// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


// Properties for `NetworkPolicyPort`.
type NetworkPolicyPortProps struct {
	// End port (relative to `port`).
	//
	// Only applies if `port` is defined.
	// Use this to specify a port range, rather that a specific one.
	EndPort *float64 `field:"optional" json:"endPort" yaml:"endPort"`
	// Specific port number.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Protocol.
	Protocol NetworkProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

