package k8s


// ResourceSliceList is a collection of ResourceSlices.
type KubeResourceSliceListV1Alpha2Props struct {
	// Items is the list of node resource capacity objects.
	Items *[]*KubeResourceSliceV1Alpha2Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

