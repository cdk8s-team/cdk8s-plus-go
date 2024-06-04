package k8s


// NamedResourcesAttribute is a combination of an attribute name and its value.
type NamedResourcesAttributeV1Alpha2 struct {
	// Name is unique identifier among all resource instances managed by the driver on the node.
	//
	// It must be a DNS subdomain.
	Name *string `field:"required" json:"name" yaml:"name"`
	// BoolValue is a true/false value.
	Bool *bool `field:"optional" json:"bool" yaml:"bool"`
	// IntValue is a 64-bit integer.
	Int *float64 `field:"optional" json:"int" yaml:"int"`
	// IntSliceValue is an array of 64-bit integers.
	IntSlice *NamedResourcesIntSliceV1Alpha2 `field:"optional" json:"intSlice" yaml:"intSlice"`
	// QuantityValue is a quantity.
	Quantity Quantity `field:"optional" json:"quantity" yaml:"quantity"`
	// StringValue is a string.
	String *string `field:"optional" json:"string" yaml:"string"`
	// StringSliceValue is an array of strings.
	StringSlice *NamedResourcesStringSliceV1Alpha2 `field:"optional" json:"stringSlice" yaml:"stringSlice"`
	// VersionValue is a semantic version according to semver.org spec 2.0.0.
	Version *string `field:"optional" json:"version" yaml:"version"`
}

