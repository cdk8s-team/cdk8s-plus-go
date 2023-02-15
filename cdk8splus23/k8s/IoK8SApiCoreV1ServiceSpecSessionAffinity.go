package k8s


// Supports "ClientIP" and "None".
//
// Used to maintain session affinity. Enable client IP based session affinity. Must be ClientIP or None. Defaults to None. More info: https://kubernetes.io/docs/concepts/services-networking/service/#virtual-ips-and-service-proxies
//
// Possible enum values:
// - `"ClientIP"` is the Client IP based.
// - `"None"` - no session affinity.
type IoK8SApiCoreV1ServiceSpecSessionAffinity string

const (
	// ClientIP.
	IoK8SApiCoreV1ServiceSpecSessionAffinity_CLIENT_IP IoK8SApiCoreV1ServiceSpecSessionAffinity = "CLIENT_IP"
	// None.
	IoK8SApiCoreV1ServiceSpecSessionAffinity_NONE IoK8SApiCoreV1ServiceSpecSessionAffinity = "NONE"
)

