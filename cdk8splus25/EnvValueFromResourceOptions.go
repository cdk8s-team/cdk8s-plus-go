package cdk8splus25


// Options to specify an environment variable value from a resource.
type EnvValueFromResourceOptions struct {
	// The container to select the value from.
	Container Container `field:"optional" json:"container" yaml:"container"`
	// The output format of the exposed resource.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

