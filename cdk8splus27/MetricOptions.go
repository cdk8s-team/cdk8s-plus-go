package cdk8splus27


// Base options for a Metric.
type MetricOptions struct {
	// The name of the metric to scale on.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The target metric value that will trigger scaling.
	Target MetricTarget `field:"required" json:"target" yaml:"target"`
	// A selector to find a metric by label.
	//
	// When set, it is passed as an additional parameter to the metrics server
	// for more specific metrics scoping.
	// Default: - Just the metric 'name' will be used to gather metrics.
	//
	LabelSelector LabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
}

