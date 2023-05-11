package cdk8splus26


// Options for `PodConnections.allowFrom`.
type PodConnectionsAllowFromOptions struct {
	// Which isolation should be applied to establish the connection.
	Isolation PodConnectionsIsolation `field:"optional" json:"isolation" yaml:"isolation"`
	// Ports to allow incoming traffic to.
	Ports *[]NetworkPolicyPort `field:"optional" json:"ports" yaml:"ports"`
}

