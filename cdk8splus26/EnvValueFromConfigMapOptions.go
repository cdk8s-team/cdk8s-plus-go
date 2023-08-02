package cdk8splus26


// Options to specify an envionment variable value from a ConfigMap key.
type EnvValueFromConfigMapOptions struct {
	// Specify whether the ConfigMap or its key must be defined.
	// Default: false.
	//
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

