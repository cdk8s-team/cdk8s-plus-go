package cdk8splus27


// Options for `WorkloadScheduling.spread`.
type WorkloadSchedulingSpreadOptions struct {
	// Which topology to spread on.
	// Default: - Topology.HOSTNAME
	//
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the spread is optional, with this weight score.
	// Default: - no weight. spread is assumed to be required.
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

