package k8s


// EvictionList contains a list of Eviction resources.
type KubeEvictionListV1Alpha1Props struct {
	// items is the list of Evictions.
	Items *[]*KubeEvictionV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// metadata is the standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

