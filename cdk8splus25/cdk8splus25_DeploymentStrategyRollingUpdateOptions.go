// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Options for `DeploymentStrategy.rollingUpdate`.
type DeploymentStrategyRollingUpdateOptions struct {
	// The maximum number of pods that can be scheduled above the desired number of pods.
	//
	// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
	// Absolute number is calculated from percentage by rounding up.
	// This can not be 0 if `maxUnavailable` is 0.
	//
	// Example: when this is set to 30%, the new ReplicaSet can be scaled up immediately when the rolling update
	// starts, such that the total number of old and new pods do not exceed 130% of desired pods.
	// Once old pods have been killed, new ReplicaSet can be scaled up further, ensuring that
	// total number of pods running at any time during the update is at most 130% of desired pods.
	MaxSurge PercentOrAbsolute `field:"optional" json:"maxSurge" yaml:"maxSurge"`
	// The maximum number of pods that can be unavailable during the update.
	//
	// Value can be an absolute number (ex: 5) or a percentage of desired pods (ex: 10%).
	// Absolute number is calculated from percentage by rounding down.
	// This can not be 0 if `maxSurge` is 0.
	//
	// Example: when this is set to 30%, the old ReplicaSet can be scaled down to 70% of desired
	// pods immediately when the rolling update starts. Once new pods are ready, old ReplicaSet can
	// be scaled down further, followed by scaling up the new ReplicaSet, ensuring that the total
	// number of pods available at all times during the update is at least 70% of desired pods.
	MaxUnavailable PercentOrAbsolute `field:"optional" json:"maxUnavailable" yaml:"maxUnavailable"`
}

