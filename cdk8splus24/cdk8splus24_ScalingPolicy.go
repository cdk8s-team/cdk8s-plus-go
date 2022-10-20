// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

type ScalingPolicy struct {
	// The type and quantity of replicas to change.
	Replicas Replicas `field:"required" json:"replicas" yaml:"replicas"`
	// The amount of time the scaling policy has to continue scaling before the target metric must be revalidated.
	//
	// Must be greater than 0 seconds and no longer than 30 minutes.
	Duration cdk8s.Duration `field:"optional" json:"duration" yaml:"duration"`
}

