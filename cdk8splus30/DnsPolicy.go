package cdk8splus30


// Pod DNS policies.
type DnsPolicy string

const (
	// Any DNS query that does not match the configured cluster domain suffix, such as "www.kubernetes.io", is forwarded to the upstream nameserver inherited from the node. Cluster administrators may have extra stub-domain and upstream DNS servers configured.
	DnsPolicy_CLUSTER_FIRST DnsPolicy = "CLUSTER_FIRST"
	// For Pods running with hostNetwork, you should explicitly set its DNS policy "ClusterFirstWithHostNet".
	DnsPolicy_CLUSTER_FIRST_WITH_HOST_NET DnsPolicy = "CLUSTER_FIRST_WITH_HOST_NET"
	// The Pod inherits the name resolution configuration from the node that the pods run on.
	DnsPolicy_DEFAULT DnsPolicy = "DEFAULT"
	// It allows a Pod to ignore DNS settings from the Kubernetes environment.
	//
	// All DNS settings are supposed to be provided using the dnsConfig
	// field in the Pod Spec.
	DnsPolicy_NONE DnsPolicy = "NONE"
)

