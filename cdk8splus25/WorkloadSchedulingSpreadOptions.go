// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Options for `WorkloadScheduling.spread`.
type WorkloadSchedulingSpreadOptions struct {
	// Which topology to spread on.
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the spread is optional, with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

