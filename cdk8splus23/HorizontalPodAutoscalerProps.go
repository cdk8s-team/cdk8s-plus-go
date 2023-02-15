// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

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
	Metrics *[]Metric `field:"optional" json:"metrics" yaml:"metrics"`
	// The minimum number of replicas that can be scaled down to.
	//
	// Can be set to 0 if the alpha feature gate `HPAScaleToZero` is enabled and
	// at least one Object or External metric is configured.
	MinReplicas *float64 `field:"optional" json:"minReplicas" yaml:"minReplicas"`
	// The scaling behavior when scaling down.
	ScaleDown *ScalingRules `field:"optional" json:"scaleDown" yaml:"scaleDown"`
	// The scaling behavior when scaling up.
	ScaleUp *ScalingRules `field:"optional" json:"scaleUp" yaml:"scaleUp"`
}

