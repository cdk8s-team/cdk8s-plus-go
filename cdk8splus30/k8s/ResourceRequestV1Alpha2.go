package k8s


// ResourceRequest is a request for resources from one particular driver.
type ResourceRequestV1Alpha2 struct {
	// NamedResources describes a request for resources with the named resources model.
	NamedResources *NamedResourcesRequestV1Alpha2 `field:"optional" json:"namedResources" yaml:"namedResources"`
	// VendorParameters are arbitrary setup parameters for the requested resource.
	//
	// They are ignored while allocating a claim.
	VendorParameters interface{} `field:"optional" json:"vendorParameters" yaml:"vendorParameters"`
}

