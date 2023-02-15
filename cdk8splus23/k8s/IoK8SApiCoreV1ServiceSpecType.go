package k8s


// type determines how the Service is exposed.
//
// Defaults to ClusterIP. Valid options are ExternalName, ClusterIP, NodePort, and LoadBalancer. "ClusterIP" allocates a cluster-internal IP address for load-balancing to endpoints. Endpoints are determined by the selector or if that is not specified, by manual construction of an Endpoints object or EndpointSlice objects. If clusterIP is "None", no virtual IP is allocated and the endpoints are published as a set of endpoints rather than a virtual IP. "NodePort" builds on ClusterIP and allocates a port on every node which routes to the same endpoints as the clusterIP. "LoadBalancer" builds on NodePort and creates an external load-balancer (if supported in the current cloud) which routes to the same endpoints as the clusterIP. "ExternalName" aliases this service to the specified externalName. Several other fields do not apply to ExternalName services. More info: https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types
//
// Possible enum values:
// - `"ClusterIP"` means a service will only be accessible inside the cluster, via the cluster IP.
// - `"ExternalName"` means a service consists of only a reference to an external name that kubedns or equivalent will return as a CNAME record, with no exposing or proxying of any pods involved.
// - `"LoadBalancer"` means a service will be exposed via an external load balancer (if the cloud provider supports it), in addition to 'NodePort' type.
// - `"NodePort"` means a service will be exposed on one port of every node, in addition to 'ClusterIP' type.
type IoK8SApiCoreV1ServiceSpecType string

const (
	// ClusterIP.
	IoK8SApiCoreV1ServiceSpecType_CLUSTER_IP IoK8SApiCoreV1ServiceSpecType = "CLUSTER_IP"
	// ExternalName.
	IoK8SApiCoreV1ServiceSpecType_EXTERNAL_NAME IoK8SApiCoreV1ServiceSpecType = "EXTERNAL_NAME"
	// LoadBalancer.
	IoK8SApiCoreV1ServiceSpecType_LOAD_BALANCER IoK8SApiCoreV1ServiceSpecType = "LOAD_BALANCER"
	// NodePort.
	IoK8SApiCoreV1ServiceSpecType_NODE_PORT IoK8SApiCoreV1ServiceSpecType = "NODE_PORT"
)

