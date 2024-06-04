package cdk8splus30


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

