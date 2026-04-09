package k8s


// ResourceClaimList is a collection of claims.
type KubeResourceClaimListProps struct {
	// Items is the list of resource claims.
	Items *[]*KubeResourceClaimProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

