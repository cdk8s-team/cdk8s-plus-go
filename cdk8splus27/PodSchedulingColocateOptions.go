package cdk8splus27


// Options for `PodScheduling.colocate`.
type PodSchedulingColocateOptions struct {
	// Which topology to coloate on.
	// Default: - Topology.HOSTNAME
	//
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the co-location is optional (soft), with this weight score.
	// Default: - no weight. co-location is assumed to be required (hard).
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

