// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Isolation determines which policies are created when allowing connections from a a pod / workload to peers.
type PodConnectionsIsolation string

const (
	// Only creates network policies that select the pod.
	PodConnectionsIsolation_POD PodConnectionsIsolation = "POD"
	// Only creates network policies that select the peer.
	PodConnectionsIsolation_PEER PodConnectionsIsolation = "PEER"
)

