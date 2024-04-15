package cdk8splus29


// Options to specify an environment variable value from the process environment.
type EnvValueFromProcessOptions struct {
	// Specify whether the key must exist in the environment.
	//
	// If this is set to true, and the key does not exist, an error will thrown.
	// Default: false.
	//
	Required *bool `field:"optional" json:"required" yaml:"required"`
}

