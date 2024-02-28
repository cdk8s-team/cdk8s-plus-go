package cdk8splus28


// Options for `StatefulSetUpdateStrategy.rollingUpdate`.
type StatefulSetUpdateStrategyRollingUpdateOptions struct {
	// If specified, all Pods with an ordinal that is greater than or equal to the partition will be updated when the StatefulSet's .spec.template is updated. All Pods with an ordinal that is less than the partition will not be updated, and, even if they are deleted, they will be recreated at the previous version.
	//
	// If the partition is greater than replicas, updates to the pod template will not be propagated to Pods.
	// In most cases you will not need to use a partition, but they are useful if you want to stage an
	// update, roll out a canary, or perform a phased roll out.
	// See: https://kubernetes.io/docs/concepts/workloads/controllers/statefulset/#partitions
	//
	// Default: 0.
	//
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
}

