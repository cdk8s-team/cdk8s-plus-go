package k8s


// podManagementPolicy controls how pods are created during initial scale up, when replacing pods on nodes, or when scaling down.
//
// The default policy is `OrderedReady`, where pods are created in increasing order (pod-0, then pod-1, etc) and the controller will wait until each pod is ready before continuing. When scaling down, the pods are removed in the opposite order. The alternative policy is `Parallel` which will create pods in parallel to match the desired scale without waiting, and on scale down will delete all pods at once.
//
// Possible enum values:
// - `"OrderedReady"` will create pods in strictly increasing order on scale up and strictly decreasing order on scale down, progressing only when the previous pod is ready or terminated. At most one pod will be changed at any time.
// - `"Parallel"` will create and delete pods as soon as the stateful set replica count is changed, and will not wait for pods to be ready or complete termination.
type IoK8SApiAppsV1StatefulSetSpecPodManagementPolicy string

const (
	// OrderedReady.
	IoK8SApiAppsV1StatefulSetSpecPodManagementPolicy_ORDERED_READY IoK8SApiAppsV1StatefulSetSpecPodManagementPolicy = "ORDERED_READY"
	// Parallel.
	IoK8SApiAppsV1StatefulSetSpecPodManagementPolicy_PARALLEL IoK8SApiAppsV1StatefulSetSpecPodManagementPolicy = "PARALLEL"
)

