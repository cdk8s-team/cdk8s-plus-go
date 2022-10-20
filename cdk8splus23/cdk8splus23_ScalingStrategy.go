// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


type ScalingStrategy string

const (
	// Use the policy that provisions the most changes.
	ScalingStrategy_MAX_CHANGE ScalingStrategy = "MAX_CHANGE"
	// Use the policy that provisions the least amount of changes.
	ScalingStrategy_MIN_CHANGE ScalingStrategy = "MIN_CHANGE"
	// Disables scaling in this direction.
	// Deprecated: - Omit the ScalingRule instead.
	ScalingStrategy_DISABLED ScalingStrategy = "DISABLED"
)

