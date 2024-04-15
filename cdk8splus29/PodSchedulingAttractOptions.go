package cdk8splus29


// Options for `PodScheduling.attract`.
type PodSchedulingAttractOptions struct {
	// Indicates the attraction is optional (soft), with this weight score.
	// Default: - no weight. assignment is assumed to be required (hard).
	//
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

