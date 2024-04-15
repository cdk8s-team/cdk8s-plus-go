package cdk8splus29


// Properties for `NetworkPolicyPort`.
type NetworkPolicyPortProps struct {
	// End port (relative to `port`).
	//
	// Only applies if `port` is defined.
	// Use this to specify a port range, rather that a specific one.
	// Default: - not a port range.
	//
	EndPort *float64 `field:"optional" json:"endPort" yaml:"endPort"`
	// Specific port number.
	// Default: - all ports are allowed.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Protocol.
	// Default: NetworkProtocol.TCP
	//
	Protocol NetworkProtocol `field:"optional" json:"protocol" yaml:"protocol"`
}

