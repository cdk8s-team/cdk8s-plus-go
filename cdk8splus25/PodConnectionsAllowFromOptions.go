package cdk8splus25


// Options for `PodConnections.allowFrom`.
type PodConnectionsAllowFromOptions struct {
	// Which isolation should be applied to establish the connection.
	// Default: - unset, isolates both the pod and the peer.
	//
	Isolation PodConnectionsIsolation `field:"optional" json:"isolation" yaml:"isolation"`
	// Ports to allow incoming traffic to.
	// Default: - The pod ports.
	//
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

