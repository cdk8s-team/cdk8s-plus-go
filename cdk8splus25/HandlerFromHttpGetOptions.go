package cdk8splus25


// Options for `Handler.fromHttpGet`.
type HandlerFromHttpGetOptions struct {
	// The TCP port to use when sending the GET request.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

