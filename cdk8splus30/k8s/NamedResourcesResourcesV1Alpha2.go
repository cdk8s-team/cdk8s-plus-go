package k8s


// NamedResourcesResources is used in ResourceModel.
type NamedResourcesResourcesV1Alpha2 struct {
	// The list of all individual resources instances currently available.
	Instances *[]*NamedResourcesInstanceV1Alpha2 `field:"required" json:"instances" yaml:"instances"`
}

