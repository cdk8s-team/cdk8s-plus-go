// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Network protocols.
type Protocol string

const (
	// TCP.
	Protocol_TCP Protocol = "TCP"
	// UDP.
	Protocol_UDP Protocol = "UDP"
	// SCTP.
	Protocol_SCTP Protocol = "SCTP"
)

