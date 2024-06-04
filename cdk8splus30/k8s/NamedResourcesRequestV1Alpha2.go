package k8s


// NamedResourcesRequest is used in ResourceRequestModel.
type NamedResourcesRequestV1Alpha2 struct {
	// Selector is a CEL expression which must evaluate to true if a resource instance is suitable.
	//
	// The language is as defined in https://kubernetes.io/docs/reference/using-api/cel/
	//
	// In addition, for each type NamedResourcesin AttributeValue there is a map that resolves to the corresponding value of the instance under evaluation. For example:
	//
	// attributes.quantity["a"].isGreaterThan(quantity("0")) &&
	// attributes.stringslice["b"].isSorted()
	Selector *string `field:"required" json:"selector" yaml:"selector"`
}

