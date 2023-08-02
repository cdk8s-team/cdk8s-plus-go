package cdk8splus27


// Options for `PodConnections.allowTo`.
type PodConnectionsAllowToOptions struct {
	// Which isolation should be applied to establish the connection.
	// Default: - unset, isolates both the pod and the peer.
	//
	Isolation PodConnectionsIsolation `field:"optional" json:"isolation" yaml:"isolation"`
	// Ports to allow outgoing traffic to.
	// Default: - If the peer is a managed pod, take its ports. Otherwise, all ports are allowed.
	//
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

