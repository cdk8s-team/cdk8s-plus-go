package k8s


// ResourceClaimTemplateList is a collection of claim templates.
type KubeResourceClaimTemplateListV1Alpha2Props struct {
	// Items is the list of resource claim templates.
	Items *[]*KubeResourceClaimTemplateV1Alpha2Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

