package k8s


// PodSchedulingContextList is a collection of Pod scheduling objects.
type KubePodSchedulingContextListV1Alpha2Props struct {
	// Items is the list of PodSchedulingContext objects.
	Items *[]*KubePodSchedulingContextV1Alpha2Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

