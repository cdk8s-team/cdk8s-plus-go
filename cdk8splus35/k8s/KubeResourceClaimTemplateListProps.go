package k8s


// ResourceClaimTemplateList is a collection of claim templates.
type KubeResourceClaimTemplateListProps struct {
	// Items is the list of resource claim templates.
	Items *[]*KubeResourceClaimTemplateProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

