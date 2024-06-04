package k8s


// ResourceClassParameters defines resource requests for a ResourceClass in an in-tree format understood by Kubernetes.
type KubeResourceClassParametersV1Alpha2Props struct {
	// Filters describes additional contraints that must be met when using the class.
	Filters *[]*ResourceFilterV1Alpha2 `field:"optional" json:"filters" yaml:"filters"`
	// If this object was created from some other resource, then this links back to that resource.
	//
	// This field is used to find the in-tree representation of the class parameters when the parameter reference of the class refers to some unknown type.
	GeneratedFrom *ResourceClassParametersReferenceV1Alpha2 `field:"optional" json:"generatedFrom" yaml:"generatedFrom"`
	// Standard object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// VendorParameters are arbitrary setup parameters for all claims using this class.
	//
	// They are ignored while allocating the claim. There must not be more than one entry per driver.
	VendorParameters *[]*VendorParametersV1Alpha2 `field:"optional" json:"vendorParameters" yaml:"vendorParameters"`
}

