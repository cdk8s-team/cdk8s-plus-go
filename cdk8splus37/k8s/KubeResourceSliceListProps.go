package k8s


// ResourceSliceList is a collection of ResourceSlices.
type KubeResourceSliceListProps struct {
	// Items is the list of resource ResourceSlices.
	Items *[]*KubeResourceSliceProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

