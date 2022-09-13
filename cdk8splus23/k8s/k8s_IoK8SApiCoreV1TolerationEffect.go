package k8s


// Effect indicates the taint effect to match.
//
// Empty means match all taint effects. When specified, allowed values are NoSchedule, PreferNoSchedule and NoExecute.
//
// Possible enum values:
// - `"NoExecute"` Evict any already-running pods that do not tolerate the taint. Currently enforced by NodeController.
// - `"NoSchedule"` Do not allow new pods to schedule onto the node unless they tolerate the taint, but allow all pods submitted to Kubelet without going through the scheduler to start, and allow all already-running pods to continue running. Enforced by the scheduler.
// - `"PreferNoSchedule"` Like TaintEffectNoSchedule, but the scheduler tries not to schedule new pods onto the node, rather than prohibiting new pods from scheduling onto the node entirely. Enforced by the scheduler.
type IoK8SApiCoreV1TolerationEffect string

const (
	// NoExecute.
	IoK8SApiCoreV1TolerationEffect_NO_EXECUTE IoK8SApiCoreV1TolerationEffect = "NO_EXECUTE"
	// NoSchedule.
	IoK8SApiCoreV1TolerationEffect_NO_SCHEDULE IoK8SApiCoreV1TolerationEffect = "NO_SCHEDULE"
	// PreferNoSchedule.
	IoK8SApiCoreV1TolerationEffect_PREFER_NO_SCHEDULE IoK8SApiCoreV1TolerationEffect = "PREFER_NO_SCHEDULE"
)

