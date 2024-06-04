package k8s


// NamedResourcesIntSlice contains a slice of 64-bit integers.
type NamedResourcesIntSliceV1Alpha2 struct {
	// Ints is the slice of 64-bit integers.
	Ints *[]*float64 `field:"required" json:"ints" yaml:"ints"`
}

