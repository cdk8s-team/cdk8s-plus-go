package cdk8splus29


// Options for `Metric.containerResource()`.
type MetricContainerResourceOptions struct {
	// Container where the metric can be found.
	Container Container `field:"required" json:"container" yaml:"container"`
	// Target metric value that will trigger scaling.
	Target MetricTarget `field:"required" json:"target" yaml:"target"`
}

