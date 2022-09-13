package k8s


// ConditionType refers to a condition in the pod's condition list with matching type.
//
// Possible enum values:
// - `"ContainersReady"` indicates whether all containers in the pod are ready.
// - `"Initialized"` means that all init containers in the pod have started successfully.
// - `"PodScheduled"` represents status of the scheduling process for this pod.
// - `"Ready"` means the pod is able to service requests and should be added to the load balancing pools of all matching services.
type IoK8SApiCoreV1PodReadinessGateConditionType string

const (
	// ContainersReady.
	IoK8SApiCoreV1PodReadinessGateConditionType_CONTAINERS_READY IoK8SApiCoreV1PodReadinessGateConditionType = "CONTAINERS_READY"
	// Initialized.
	IoK8SApiCoreV1PodReadinessGateConditionType_INITIALIZED IoK8SApiCoreV1PodReadinessGateConditionType = "INITIALIZED"
	// PodScheduled.
	IoK8SApiCoreV1PodReadinessGateConditionType_POD_SCHEDULED IoK8SApiCoreV1PodReadinessGateConditionType = "POD_SCHEDULED"
	// Ready.
	IoK8SApiCoreV1PodReadinessGateConditionType_READY IoK8SApiCoreV1PodReadinessGateConditionType = "READY"
)

