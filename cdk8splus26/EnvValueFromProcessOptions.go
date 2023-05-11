package cdk8splus26


// Options to specify an environment variable value from the process environment.
type EnvValueFromProcessOptions struct {
	// Specify whether the key must exist in the environment.
	//
	// If this is set to true, and the key does not exist, an error will thrown.
	Required *bool `field:"optional" json:"required" yaml:"required"`
}

