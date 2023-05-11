package cdk8splus26


// Options for `PodScheduling.separate`.
type PodSchedulingSeparateOptions struct {
	// Which topology to separate on.
	Topology Topology `field:"optional" json:"topology" yaml:"topology"`
	// Indicates the separation is optional (soft), with this weight score.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

