package k8s


// PodGroupTemplate represents a template for a set of pods with a scheduling policy.
type PodGroupTemplateV1Alpha3 struct {
	// name is a unique identifier for the PodGroupTemplate within the Workload.
	//
	// It must be a DNS label. This field is immutable.
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedulingPolicy defines the scheduling policy for this PodGroupTemplate.
	SchedulingPolicy *PodGroupSchedulingPolicyV1Alpha3 `field:"required" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// disruptionMode defines the mode in which a given PodGroup can be disrupted.
	//
	// One of Single, All. This field is immutable.
	DisruptionMode *DisruptionModeV1Alpha3 `field:"optional" json:"disruptionMode" yaml:"disruptionMode"`
	// preemptionPolicy is the Policy for preempting pods/podgroups with lower priority.
	//
	// One of Never, PreemptLowerPriority. This field is immutable. This field is available only when the PodGroupPreemptionPolicy feature gate is enabled.
	PreemptionPolicy *string `field:"optional" json:"preemptionPolicy" yaml:"preemptionPolicy"`
	// priority is the value of priority of pod groups created from this template.
	//
	// Various system components use this field to find the priority of the pod group. The higher the value, the higher the priority. This field is immutable.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// priorityClassName indicates the priority that should be considered when scheduling a pod group created from this template.
	//
	// This field is immutable.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// resourceClaims defines which ResourceClaims may be shared among Pods in the group.
	//
	// Pods consume the devices allocated to a PodGroup's claim by defining a claim in its own Spec.ResourceClaims that matches the PodGroup's claim exactly. The claim must have the same name and refer to the same ResourceClaim or ResourceClaimTemplate.
	//
	// This is a beta-level field and requires that the DRAWorkloadResourceClaims feature gate is enabled.
	//
	// This field is immutable.
	ResourceClaims *[]*PodGroupResourceClaimV1Alpha3 `field:"optional" json:"resourceClaims" yaml:"resourceClaims"`
	// schedulingConstraints defines optional scheduling constraints (e.g. topology) for this PodGroupTemplate. This field is only available when the TopologyAwareWorkloadScheduling feature gate is enabled. This field is immutable.
	SchedulingConstraints *PodGroupSchedulingConstraintsV1Alpha3 `field:"optional" json:"schedulingConstraints" yaml:"schedulingConstraints"`
}

