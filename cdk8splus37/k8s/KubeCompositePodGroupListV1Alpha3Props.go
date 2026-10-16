package k8s


// CompositePodGroupList contains a list of CompositePodGroup resources.
type KubeCompositePodGroupListV1Alpha3Props struct {
	// Items is the list of CompositePodGroups.
	Items *[]*KubeCompositePodGroupV1Alpha3Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

