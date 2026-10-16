package k8s


// EvictionRequestTarget contains a reference to an object that should be evicted.
type EvictionRequestTargetV1Alpha1 struct {
	// pod references a pod that is subject to eviction/termination.
	//
	// Pods that are part of a PodGroup (.spec.schedulingGroup is set) are not supported.
	Pod *EvictionRequestPodReferenceV1Alpha1 `field:"optional" json:"pod" yaml:"pod"`
}

