package cdk8splus26


// Options for `PodScheduling.colocate`.
type PodSchedulingColocateOptions struct {
	// Which topology to coloate on.
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the co-location is optional (soft), with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

