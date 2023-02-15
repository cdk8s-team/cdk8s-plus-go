package k8s


// Protocol for port. Must be UDP, TCP, or SCTP. Defaults to "TCP".
//
// Possible enum values:
// - `"SCTP"` is the SCTP protocol.
// - `"TCP"` is the TCP protocol.
// - `"UDP"` is the UDP protocol.
type IoK8SApiCoreV1ContainerPortProtocol string

const (
	// SCTP.
	IoK8SApiCoreV1ContainerPortProtocol_SCTP IoK8SApiCoreV1ContainerPortProtocol = "SCTP"
	// TCP.
	IoK8SApiCoreV1ContainerPortProtocol_TCP IoK8SApiCoreV1ContainerPortProtocol = "TCP"
	// UDP.
	IoK8SApiCoreV1ContainerPortProtocol_UDP IoK8SApiCoreV1ContainerPortProtocol = "UDP"
)

