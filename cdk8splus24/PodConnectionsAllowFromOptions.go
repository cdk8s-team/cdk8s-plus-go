// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Options for `PodConnections.allowFrom`.
type PodConnectionsAllowFromOptions struct {
	// Which isolation should be applied to establish the connection.
	Isolation PodConnectionsIsolation `field:"optional" json:"isolation" yaml:"isolation"`
	// Ports to allow incoming traffic to.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

