package cdk8splus31


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

