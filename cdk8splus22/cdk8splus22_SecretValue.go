// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Represents a specific value in JSON secret.
type SecretValue struct {
	// The JSON key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The secret.
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
}

