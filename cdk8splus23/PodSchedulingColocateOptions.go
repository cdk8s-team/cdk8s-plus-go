// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Options for `PodScheduling.colocate`.
type PodSchedulingColocateOptions struct {
	// Which topology to coloate on.
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the co-location is optional (soft), with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

