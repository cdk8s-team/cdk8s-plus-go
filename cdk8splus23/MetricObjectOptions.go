// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Options for `Metric.object()`.
type MetricObjectOptions struct {
	// The name of the metric to scale on.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The target metric value that will trigger scaling.
	Target MetricTarget `field:"required" json:"target" yaml:"target"`
	// A selector to find a metric by label.
	//
	// When set, it is passed as an additional parameter to the metrics server
	// for more specific metrics scoping.
	LabelSelector LabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// Resource where the metric can be found.
	Object IResource `field:"required" json:"object" yaml:"object"`
}

