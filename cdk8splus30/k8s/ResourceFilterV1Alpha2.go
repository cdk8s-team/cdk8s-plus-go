package k8s


// ResourceFilter is a filter for resources from one particular driver.
type ResourceFilterV1Alpha2 struct {
	// DriverName is the name used by the DRA driver kubelet plugin.
	DriverName *string `field:"optional" json:"driverName" yaml:"driverName"`
	// NamedResources describes a resource filter using the named resources model.
	NamedResources *NamedResourcesFilterV1Alpha2 `field:"optional" json:"namedResources" yaml:"namedResources"`
}

