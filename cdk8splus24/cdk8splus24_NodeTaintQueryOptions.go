// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Options for `NodeTaintQuery`.
type NodeTaintQueryOptions struct {
	// The taint effect to match.
	Effect TaintEffect `field:"optional" json:"effect" yaml:"effect"`
	// How much time should a pod that tolerates the `NO_EXECUTE` effect be bound to the node.
	//
	// Only applies for the `NO_EXECUTE` effect.
	EvictAfter cdk8s.Duration `field:"optional" json:"evictAfter" yaml:"evictAfter"`
}

