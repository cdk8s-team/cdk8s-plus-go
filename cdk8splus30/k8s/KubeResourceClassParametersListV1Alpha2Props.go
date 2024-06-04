package k8s


// ResourceClassParametersList is a collection of ResourceClassParameters.
type KubeResourceClassParametersListV1Alpha2Props struct {
	// Items is the list of node resource capacity objects.
	Items *[]*KubeResourceClassParametersV1Alpha2Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

