// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


// Options to specify an environment variable value from a resource.
type EnvValueFromResourceOptions struct {
	// The container to select the value from.
	Container Container `field:"optional" json:"container" yaml:"container"`
	// The output format of the exposed resource.
	Divisor *string `field:"optional" json:"divisor" yaml:"divisor"`
}

