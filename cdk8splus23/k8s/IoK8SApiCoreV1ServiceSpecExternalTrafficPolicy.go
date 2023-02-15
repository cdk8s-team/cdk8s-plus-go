package k8s


// externalTrafficPolicy denotes if this Service desires to route external traffic to node-local or cluster-wide endpoints.
//
// "Local" preserves the client source IP and avoids a second hop for LoadBalancer and Nodeport type services, but risks potentially imbalanced traffic spreading. "Cluster" obscures the client source IP and may cause a second hop to another node, but should have good overall load-spreading.
//
// Possible enum values:
// - `"Cluster"` specifies node-global (legacy) behavior.
// - `"Local"` specifies node-local endpoints behavior.
type IoK8SApiCoreV1ServiceSpecExternalTrafficPolicy string

const (
	// Cluster.
	IoK8SApiCoreV1ServiceSpecExternalTrafficPolicy_CLUSTER IoK8SApiCoreV1ServiceSpecExternalTrafficPolicy = "CLUSTER"
	// Local.
	IoK8SApiCoreV1ServiceSpecExternalTrafficPolicy_LOCAL IoK8SApiCoreV1ServiceSpecExternalTrafficPolicy = "LOCAL"
)

