package k8s


// EvictionSpec is a specification of an Eviction.
type EvictionSpecV1Alpha1 struct {
	// target contains a reference to an object (e.g. a pod) that should be evicted. This field is required and immutable.
	Target *EvictionTargetV1Alpha1 `field:"required" json:"target" yaml:"target"`
}

