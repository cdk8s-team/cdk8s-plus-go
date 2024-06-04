package k8s


// NamedResourcesStringSlice contains a slice of strings.
type NamedResourcesStringSliceV1Alpha2 struct {
	// Strings is the slice of strings.
	Strings *[]*string `field:"required" json:"strings" yaml:"strings"`
}

