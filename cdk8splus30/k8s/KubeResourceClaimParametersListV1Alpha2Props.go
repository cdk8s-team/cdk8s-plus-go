package k8s


// ResourceClaimParametersList is a collection of ResourceClaimParameters.
type KubeResourceClaimParametersListV1Alpha2Props struct {
	// Items is the list of node resource capacity objects.
	Items *[]*KubeResourceClaimParametersV1Alpha2Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

