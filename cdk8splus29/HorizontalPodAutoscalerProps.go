package cdk8splus29

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for HorizontalPodAutoscaler.
type HorizontalPodAutoscalerProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// The maximum number of replicas that can be scaled up to.
	MaxReplicas *float64 `field:"required" json:"maxReplicas" yaml:"maxReplicas"`
	// The workload to scale up or down.
	//
	// Scalable workload types:
	// * Deployment
	// * StatefulSet.
	Target IScalable `field:"required" json:"target" yaml:"target"`
	// The metric conditions that trigger a scale up or scale down.
	// Default: - If metrics are not provided, then the target resource
	// constraints (e.g. cpu limit) will be used as scaling metrics.
	//
	Metrics *[]Metric `field:"optional" json:"metrics" yaml:"metrics"`
	// The minimum number of replicas that can be scaled down to.
	//
	// Can be set to 0 if the alpha feature gate `HPAScaleToZero` is enabled and
	// at least one Object or External metric is configured.
	// Default: 1.
	//
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// The scaling behavior when scaling down.
	// Default: - Scale down to minReplica count with a 5 minute stabilization window.
	//
	ScaleDown *ScalingRules `field:"optional" json:"scaleDown" yaml:"scaleDown"`
	// The scaling behavior when scaling up.
	// Default: - Is the higher of:
	// * Increase no more than 4 pods per 60 seconds
	// * Double the number of pods per 60 seconds.
	//
	ScaleUp *ScalingRules `field:"optional" json:"scaleUp" yaml:"scaleUp"`
}

