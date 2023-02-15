package k8s


// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
//
// Possible enum values:
// - `"SCTP"` is the SCTP protocol.
// - `"TCP"` is the TCP protocol.
// - `"UDP"` is the UDP protocol.
type IoK8SApiCoreV1EndpointPortProtocol string

const (
	// SCTP.
	IoK8SApiCoreV1EndpointPortProtocol_SCTP IoK8SApiCoreV1EndpointPortProtocol = "SCTP"
	// TCP.
	IoK8SApiCoreV1EndpointPortProtocol_TCP IoK8SApiCoreV1EndpointPortProtocol = "TCP"
	// UDP.
	IoK8SApiCoreV1EndpointPortProtocol_UDP IoK8SApiCoreV1EndpointPortProtocol = "UDP"
)

