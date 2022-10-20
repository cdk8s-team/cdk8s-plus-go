// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Defines the scaling behavior for one direction.
type ScalingRules struct {
	// The scaling policies.
	Policies *[]*ScalingPolicy `field:"optional" json:"policies" yaml:"policies"`
	// Defines the window of past metrics that the autoscaler should consider when calculating wether or not autoscaling should occur.
	//
	// Minimum duration is 1 second, max is 1 hour.
	//
	// Example:
	//   stabilizationWindow: Duration.minutes(30)
	//   // Autoscaler considers the last 30 minutes of metrics when deciding whether to scale.
	//
	StabilizationWindow cdk8s.Duration `field:"optional" json:"stabilizationWindow" yaml:"stabilizationWindow"`
	// The strategy to use when scaling.
	Strategy ScalingStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

