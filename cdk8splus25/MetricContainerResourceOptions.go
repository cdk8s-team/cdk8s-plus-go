// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Options for `Metric.containerResource()`.
type MetricContainerResourceOptions struct {
	// Container where the metric can be found.
	Container Container `field:"required" json:"container" yaml:"container"`
	// Target metric value that will trigger scaling.
	Target MetricTarget `field:"required" json:"target" yaml:"target"`
}

