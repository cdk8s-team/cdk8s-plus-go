package k8s


// PodGroupList contains a list of PodGroup resources.
type KubePodGroupListV1Alpha2Props struct {
	// Items is the list of PodGroups.
	Items *[]*KubePodGroupV1Alpha2Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

