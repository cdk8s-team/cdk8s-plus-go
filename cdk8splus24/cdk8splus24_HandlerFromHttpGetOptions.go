// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Options for `Handler.fromHttpGet`.
type HandlerFromHttpGetOptions struct {
	// The TCP port to use when sending the GET request.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

