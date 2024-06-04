package cdk8splus30

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Options for `NodeTaintQuery`.
type NodeTaintQueryOptions struct {
	// The taint effect to match.
	// Default: - all effects are matched.
	//
	Effect TaintEffect `field:"optional" json:"effect" yaml:"effect"`
	// How much time should a pod that tolerates the `NO_EXECUTE` effect be bound to the node.
	//
	// Only applies for the `NO_EXECUTE` effect.
	// Default: - bound forever.
	//
	EvictAfter cdk8s.Duration `field:"optional" json:"evictAfter" yaml:"evictAfter"`
}

