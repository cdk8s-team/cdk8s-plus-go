// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Options for `PodScheduling.attract`.
type PodSchedulingAttractOptions struct {
	// Indicates the attraction is optional (soft), with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

