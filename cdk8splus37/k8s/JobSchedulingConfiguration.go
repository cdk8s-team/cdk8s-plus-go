package k8s


// JobSchedulingConfiguration composes the reusable workload-aware scheduling building blocks.
type JobSchedulingConfiguration struct {
	// DisruptionMode defines the mode in which the Job's pods can be disrupted.
	//
	// One of Single, All. This field is immutable after creation: it may not be added or removed, and the selected mode may not be changed.
	DisruptionMode *WorkloadPodGroupDisruptionModeV1Alpha3 `field:"optional" json:"disruptionMode" yaml:"disruptionMode"`
	// ResourceClaims defines which ResourceClaims may be shared among Pods in the Job.
	//
	// Pods consume the devices allocated to a PodGroup's claim by defining a claim in its own Spec.ResourceClaims that matches the PodGroup's claim exactly. The claim must have the same name and refer to the same ResourceClaim or ResourceClaimTemplate. At most 4 claims may be set, matching the limit on the resulting PodGroup. This list is immutable after creation: entries may neither be added, removed, nor modified.
	ResourceClaims *[]*WorkloadPodGroupResourceClaimV1Alpha3 `field:"optional" json:"resourceClaims" yaml:"resourceClaims"`
	// SchedulingConstraints defines scheduling constraints (e.g. topology) for the Job's pods. This field is immutable after creation.
	SchedulingConstraints *WorkloadPodGroupSchedulingConstraintsV1Alpha3 `field:"optional" json:"schedulingConstraints" yaml:"schedulingConstraints"`
	// SchedulingPolicy defines the scheduling policy for this Job.
	//
	// Exactly one of Basic or Gang must be set. This field is immutable after creation: the policy may not be added or removed. The policy variant (basic/gang) is frozen by hand-written validation; only schedulingPolicy.gang.minCount may be changed.
	SchedulingPolicy *WorkloadPodGroupSchedulingPolicyV1Alpha3 `field:"optional" json:"schedulingPolicy" yaml:"schedulingPolicy"`
}

