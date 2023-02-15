package k8s


// Required.
//
// The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
//
// Possible enum values:
// - `"NoExecute"` Evict any already-running pods that do not tolerate the taint. Currently enforced by NodeController.
// - `"NoSchedule"` Do not allow new pods to schedule onto the node unless they tolerate the taint, but allow all pods submitted to Kubelet without going through the scheduler to start, and allow all already-running pods to continue running. Enforced by the scheduler.
// - `"PreferNoSchedule"` Like TaintEffectNoSchedule, but the scheduler tries not to schedule new pods onto the node, rather than prohibiting new pods from scheduling onto the node entirely. Enforced by the scheduler.
type IoK8SApiCoreV1TaintEffect string

const (
	// NoExecute.
	IoK8SApiCoreV1TaintEffect_NO_EXECUTE IoK8SApiCoreV1TaintEffect = "NO_EXECUTE"
	// NoSchedule.
	IoK8SApiCoreV1TaintEffect_NO_SCHEDULE IoK8SApiCoreV1TaintEffect = "NO_SCHEDULE"
	// PreferNoSchedule.
	IoK8SApiCoreV1TaintEffect_PREFER_NO_SCHEDULE IoK8SApiCoreV1TaintEffect = "PREFER_NO_SCHEDULE"
)

