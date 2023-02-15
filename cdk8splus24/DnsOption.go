// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Custom DNS option.
type DnsOption struct {
	// Option name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Option value.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

