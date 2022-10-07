// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Options for `StatefulSetUpdateStrategy.rollingUpdate`.
type StatefulSetUpdateStrategyRollingUpdateOptions struct {
	// If specified, all Pods with an ordinal that is greater than or equal to the partition will be updated when the StatefulSet's .spec.template is updated. All Pods with an ordinal that is less than the partition will not be updated, and, even if they are deleted, they will be recreated at the previous version.
	//
	// If the partition is greater than replicas, updates to the pod template will not be propagated to Pods.
	// In most cases you will not need to use a partition, but they are useful if you want to stage an
	// update, roll out a canary, or perform a phased roll out.
	// See: https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#partitions
	//
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
}

