package k8s


// EvictionPodReference contains enough information to locate the referenced pod inside the same namespace.
type EvictionPodReferenceV1Alpha1 struct {
	// name of the target.
	//
	// This field is required.
	Name *string `field:"required" json:"name" yaml:"name"`
	// uid of the target.
	//
	// It can be found in .metadata.uid of the target and is a lowercase UUID in 8-4-4-4-12 format. This field is required.
	Uid *string `field:"required" json:"uid" yaml:"uid"`
}

