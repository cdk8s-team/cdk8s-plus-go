package k8s


// PodReadinessGate contains the reference to a pod condition.
type PodReadinessGate struct {
	// ConditionType refers to a condition in the pod's condition list with matching type.
	//
	// Possible enum values:
	// - `"ContainersReady"` indicates whether all containers in the pod are ready.
	// - `"Initialized"` means that all init containers in the pod have started successfully.
	// - `"PodScheduled"` represents status of the scheduling process for this pod.
	// - `"Ready"` means the pod is able to service requests and should be added to the load balancing pools of all matching services.
	ConditionType IoK8SApiCoreV1PodReadinessGateConditionType `field:"required" json:"conditionType" yaml:"conditionType"`
}

