package k8s


// DeviceAttribute must have exactly one field set.
type DeviceAttributeV1Beta2 struct {
	// BoolValue is a true/false value.
	Bool *bool `field:"optional" json:"bool" yaml:"bool"`
	// BoolValues is a non-empty list of true/false values.
	Bools *[]*bool `field:"optional" json:"bools" yaml:"bools"`
	// IntValue is a number.
	Int *float64 `field:"optional" json:"int" yaml:"int"`
	// IntValues is a non-empty list of numbers.
	//
	// This is an alpha field and requires enabling the DRAListTypeAttributes feature gate.
	Ints *[]*float64 `field:"optional" json:"ints" yaml:"ints"`
	// StringValue is a string.
	//
	// Must not be longer than 64 characters.
	String *string `field:"optional" json:"string" yaml:"string"`
	// StringValues is a non-empty list of strings. Each string must not be longer than 64 characters.
	//
	// This is an alpha field and requires enabling the DRAListTypeAttributes feature gate.
	Strings *[]*string `field:"optional" json:"strings" yaml:"strings"`
	// VersionValue is a semantic version according to semver.org spec 2.0.0. Must not be longer than 64 characters.
	Version *string `field:"optional" json:"version" yaml:"version"`
	// VersionValues is a non-empty list of semantic versions according to semver.org spec 2.0.0. Each version string must not be longer than 64 characters.
	//
	// This is an alpha field and requires enabling the DRAListTypeAttributes feature gate.
	Versions *[]*string `field:"optional" json:"versions" yaml:"versions"`
}

