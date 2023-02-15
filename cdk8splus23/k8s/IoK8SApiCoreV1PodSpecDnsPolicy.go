package k8s


// Set DNS policy for the pod.
//
// Defaults to "ClusterFirst". Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'. DNS parameters given in DNSConfig will be merged with the policy selected with DNSPolicy. To have DNS options set along with hostNetwork, you have to specify DNS policy explicitly to 'ClusterFirstWithHostNet'.
//
// Possible enum values:
// - `"ClusterFirst"` indicates that the pod should use cluster DNS first unless hostNetwork is true, if it is available, then fall back on the default (as determined by kubelet) DNS settings.
// - `"ClusterFirstWithHostNet"` indicates that the pod should use cluster DNS first, if it is available, then fall back on the default (as determined by kubelet) DNS settings.
// - `"Default"` indicates that the pod should use the default (as determined by kubelet) DNS settings.
// - `"None"` indicates that the pod should use empty DNS settings. DNS parameters such as nameservers and search paths should be defined via DNSConfig.
type IoK8SApiCoreV1PodSpecDnsPolicy string

const (
	// ClusterFirst.
	IoK8SApiCoreV1PodSpecDnsPolicy_CLUSTER_FIRST IoK8SApiCoreV1PodSpecDnsPolicy = "CLUSTER_FIRST"
	// ClusterFirstWithHostNet.
	IoK8SApiCoreV1PodSpecDnsPolicy_CLUSTER_FIRST_WITH_HOST_NET IoK8SApiCoreV1PodSpecDnsPolicy = "CLUSTER_FIRST_WITH_HOST_NET"
	// Default.
	IoK8SApiCoreV1PodSpecDnsPolicy_DEFAULT IoK8SApiCoreV1PodSpecDnsPolicy = "DEFAULT"
	// None.
	IoK8SApiCoreV1PodSpecDnsPolicy_NONE IoK8SApiCoreV1PodSpecDnsPolicy = "NONE"
)

