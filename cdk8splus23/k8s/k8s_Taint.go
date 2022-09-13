package k8s

import (
	"time"
)

// The node this Taint is attached to has the "effect" on any pod that does not tolerate the Taint.
type Taint struct {
	// Required.
	//
	// The effect of the taint on pods that do not tolerate the taint. Valid effects are NoSchedule, PreferNoSchedule and NoExecute.
	//
	// Possible enum values:
	// - `"NoExecute"` Evict any already-running pods that do not tolerate the taint. Currently enforced by NodeController.
	// - `"NoSchedule"` Do not allow new pods to schedule onto the node unless they tolerate the taint, but allow all pods submitted to Kubelet without going through the scheduler to start, and allow all already-running pods to continue running. Enforced by the scheduler.
	// - `"PreferNoSchedule"` Like TaintEffectNoSchedule, but the scheduler tries not to schedule new pods onto the node, rather than prohibiting new pods from scheduling onto the node entirely. Enforced by the scheduler.
	Effect IoK8SApiCoreV1TaintEffect `field:"required" json:"effect" yaml:"effect"`
	// Required.
	//
	// The taint key to be applied to a node.
	Key *string `field:"required" json:"key" yaml:"key"`
	// TimeAdded represents the time at which the taint was added.
	//
	// It is only written for NoExecute taints.
	TimeAdded *time.Time `field:"optional" json:"timeAdded" yaml:"timeAdded"`
	// The taint value corresponding to the taint key.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

