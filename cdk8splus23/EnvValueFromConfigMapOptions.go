// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Options to specify an envionment variable value from a ConfigMap key.
type EnvValueFromConfigMapOptions struct {
	// Specify whether the ConfigMap or its key must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

