// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Network protocols.
type NetworkProtocol string

const (
	// TCP.
	NetworkProtocol_TCP NetworkProtocol = "TCP"
	// UDP.
	NetworkProtocol_UDP NetworkProtocol = "UDP"
	// SCTP.
	NetworkProtocol_SCTP NetworkProtocol = "SCTP"
)

