package cdk8splus28


// Custom DNS option.
type DnsOption struct {
	// Option name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Option value.
	// Default: - No value.
	//
	Value *string `field:"optional" json:"value" yaml:"value"`
}

