// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Options for `PodScheduling.separate`.
type PodSchedulingSeparateOptions struct {
	// Which topology to separate on.
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the separation is optional (soft), with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

