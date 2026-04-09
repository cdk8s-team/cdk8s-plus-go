package cdk8splus34


type HttpHeader struct {
	// The HTTP Header name to be used.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The HTTP header value to be set.
	Value *string `field:"required" json:"value" yaml:"value"`
}

