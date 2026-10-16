package k8s


// NodePodPreemptionPolicy defines the node-level policies governing preemption for pods on this node.
type NodePodPreemptionPolicy struct {
	// DisableResizePreemption lists the owners (e.g., autoscalers, operators, administrators) that have requested to disable scheduler and Kubelet preemption for in-place pod resize on this node. If this list is non-empty, resize-induced preemption is disabled on this node. This is an alpha field and requires enabling the InPlacePodVerticalScalingSchedulerPreemption feature gate.
	DisableResizePreemption *[]*string `field:"optional" json:"disableResizePreemption" yaml:"disableResizePreemption"`
}

