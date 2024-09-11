package cdk8splus31


// Options to specify an environment variable value from a Secret.
type EnvValueFromSecretOptions struct {
	// Specify whether the Secret or its key must be defined.
	// Default: false.
	//
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

