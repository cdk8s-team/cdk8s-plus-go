package cdk8splus29


// Options for `PodScheduling.separate`.
type PodSchedulingSeparateOptions struct {
	// Which topology to separate on.
	// Default: - Topology.HOSTNAME
	//
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the separation is optional (soft), with this weight score.
	// Default: - no weight. separation is assumed to be required (hard).
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

