package cdk8splus28


// Options for `Handler.fromHttpGet`.
type HandlerFromHttpGetOptions struct {
	// The TCP port to use when sending the GET request.
	// Default: - defaults to `container.port`.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

