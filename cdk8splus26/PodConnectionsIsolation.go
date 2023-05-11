package cdk8splus26


// Isolation determines which policies are created when allowing connections from a a pod / workload to peers.
type PodConnectionsIsolation string

const (
	// Only creates network policies that select the pod.
	PodConnectionsIsolation_POD PodConnectionsIsolation = "POD"
	// Only creates network policies that select the peer.
	PodConnectionsIsolation_PEER PodConnectionsIsolation = "PEER"
)

