package k8s


// EvictionRequestList contains a list of EvictionRequests resources.
type KubeEvictionRequestListV1Alpha1Props struct {
	// items is the list of EvictionRequests.
	Items *[]*KubeEvictionRequestV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// metadata is the standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

