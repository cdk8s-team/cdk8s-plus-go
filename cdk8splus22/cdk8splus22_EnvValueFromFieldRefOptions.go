// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Options to specify an environment variable value from a field reference.
type EnvValueFromFieldRefOptions struct {
	// Version of the schema the FieldPath is written in terms of.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// The key to select the pod label or annotation.
	Key *string `field:"optional" json:"key" yaml:"key"`
}

