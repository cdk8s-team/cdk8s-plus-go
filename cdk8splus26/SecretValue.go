// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


// Represents a specific value in JSON secret.
type SecretValue struct {
	// The JSON key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The secret.
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
}

