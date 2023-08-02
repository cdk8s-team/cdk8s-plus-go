package cdk8splus26

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Defines the scaling behavior for one direction.
type ScalingRules struct {
	// The scaling policies.
	// Default: * Scale up
	//   * Increase no more than 4 pods per 60 seconds
	//   * Double the number of pods per 60 seconds
	// * Scale down
	// * Decrease to minReplica count.
	//
	Policies *[]*ScalingPolicy `field:"optional" json:"policies" yaml:"policies"`
	// Defines the window of past metrics that the autoscaler should consider when calculating wether or not autoscaling should occur.
	//
	// Minimum duration is 1 second, max is 1 hour.
	//
	// Example:
	//   stabilizationWindow: Duration.minutes(30)
	//   // Autoscaler considers the last 30 minutes of metrics when deciding whether to scale.
	//
	// Default: * On scale down no stabilization is performed.
	// * On scale up stabilization is performed for 5 minutes.
	//
	StabilizationWindow cdk8s.Duration `field:"optional" json:"stabilizationWindow" yaml:"stabilizationWindow"`
	// The strategy to use when scaling.
	// Default: MAX_CHANGE.
	//
	Strategy ScalingStrategy `field:"optional" json:"strategy" yaml:"strategy"`
}

