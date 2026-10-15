package k8s


// Storage version of a specific resource.
type KubeStorageVersionV1Alpha1Props struct {
	// metadata is the standard object metadata.
	//
	// The name is <group>.<resource>.
	Metadata *ObjectMeta `field:"required" json:"metadata" yaml:"metadata"`
	// spec is an empty spec.
	//
	// It is here to comply with Kubernetes API style.
	Spec interface{} `field:"optional" json:"spec" yaml:"spec"`
}

