package k8s


// ResourcePoolStatusRequestList is a collection of ResourcePoolStatusRequests.
type KubeResourcePoolStatusRequestListV1Alpha3Props struct {
	// Items is the list of ResourcePoolStatusRequests.
	Items *[]*KubeResourcePoolStatusRequestV1Alpha3Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

