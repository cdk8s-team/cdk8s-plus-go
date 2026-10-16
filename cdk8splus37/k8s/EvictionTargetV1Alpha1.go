package k8s


// EvictionTarget contains a reference to an object that should be evicted.
type EvictionTargetV1Alpha1 struct {
	// pod references a pod that is subject to eviction/termination.
	//
	// Pods that are part of a PodGroup (.spec.schedulingGroup is set) are not supported.
	Pod *EvictionPodReferenceV1Alpha1 `field:"optional" json:"pod" yaml:"pod"`
}

