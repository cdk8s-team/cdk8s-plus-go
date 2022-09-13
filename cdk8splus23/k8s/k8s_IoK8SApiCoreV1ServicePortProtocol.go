package k8s


// The IP protocol for this port. Supports "TCP", "UDP", and "SCTP". Default is TCP.
//
// Possible enum values:
// - `"SCTP"` is the SCTP protocol.
// - `"TCP"` is the TCP protocol.
// - `"UDP"` is the UDP protocol.
type IoK8SApiCoreV1ServicePortProtocol string

const (
	// SCTP.
	IoK8SApiCoreV1ServicePortProtocol_SCTP IoK8SApiCoreV1ServicePortProtocol = "SCTP"
	// TCP.
	IoK8SApiCoreV1ServicePortProtocol_TCP IoK8SApiCoreV1ServicePortProtocol = "TCP"
	// UDP.
	IoK8SApiCoreV1ServicePortProtocol_UDP IoK8SApiCoreV1ServicePortProtocol = "UDP"
)

