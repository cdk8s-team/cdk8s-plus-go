package cdk8splus30


// Taint effects.
type TaintEffect string

const (
	// This means that no pod will be able to schedule onto the node unless it has a matching toleration.
	TaintEffect_NO_SCHEDULE TaintEffect = "NO_SCHEDULE"
	// This is a "preference" or "soft" version of `NO_SCHEDULE` -- the system will try to avoid placing a pod that does not tolerate the taint on the node, but it is not required.
	TaintEffect_PREFER_NO_SCHEDULE TaintEffect = "PREFER_NO_SCHEDULE"
	// This affects pods that are already running on the node as follows:.
	//
	// - Pods that do not tolerate the taint are evicted immediately.
	// - Pods that tolerate the taint without specifying `duration` remain bound forever.
	// - Pods that tolerate the taint with a specified `duration` remain bound for
	//    the specified amount of time.
	TaintEffect_NO_EXECUTE TaintEffect = "NO_EXECUTE"
)

