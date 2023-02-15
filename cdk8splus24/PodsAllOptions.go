// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Options for `Pods.all`.
type PodsAllOptions struct {
	// Namespaces the pods are allowed to be in.
	//
	// Use `Namespaces.all()` to allow all namespaces.
	Namespaces Namespaces `field:"optional" json:"namespaces" yaml:"namespaces"`
}

