package k8s


// Restart policy for all containers within the pod. One of Always, OnFailure, Never. Default to Always. More info: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#restart-policy.
//
// Possible enum values:
// - `"Always"`
// - `"Never"`
// - `"OnFailure"`.
type IoK8SApiCoreV1PodSpecRestartPolicy string

const (
	// Always.
	IoK8SApiCoreV1PodSpecRestartPolicy_ALWAYS IoK8SApiCoreV1PodSpecRestartPolicy = "ALWAYS"
	// Never.
	IoK8SApiCoreV1PodSpecRestartPolicy_NEVER IoK8SApiCoreV1PodSpecRestartPolicy = "NEVER"
	// OnFailure.
	IoK8SApiCoreV1PodSpecRestartPolicy_ON_FAILURE IoK8SApiCoreV1PodSpecRestartPolicy = "ON_FAILURE"
)

