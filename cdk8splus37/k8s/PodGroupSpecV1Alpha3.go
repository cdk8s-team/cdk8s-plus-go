package k8s


// PodGroupSpec defines the desired state of a PodGroup.
type PodGroupSpecV1Alpha3 struct {
	// schedulingPolicy defines the scheduling policy for this instance of the PodGroup.
	//
	// Controllers are expected to fill this field by copying it from a PodGroupTemplate.
	SchedulingPolicy *PodGroupSchedulingPolicyV1Alpha3 `field:"required" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// disruptionMode defines the mode in which a given PodGroup can be disrupted.
	//
	// Controllers are expected to fill this field by copying it from a PodGroupTemplate. One of Single, All. Defaults to Single if unset. This field is immutable.
	// Default: Single if unset. This field is immutable.
	//
	DisruptionMode *DisruptionModeV1Alpha3 `field:"optional" json:"disruptionMode" yaml:"disruptionMode"`
	// parentCompositePodGroupName contains the name of the parent composite pod group within the same namespace as this pod group.
	//
	// If it's nil, then this pod group is a root of a workload's hierarchy. This field is used only when the CompositePodGroup feature gate is enabled. This field is immutable.
	ParentCompositePodGroupName *string `field:"optional" json:"parentCompositePodGroupName" yaml:"parentCompositePodGroupName"`
	// preemptionPolicy is the Policy for preempting pods/podgroups with lower priority.
	//
	// One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. When Priority Admission Controller is enabled, it populates this field from PriorityClassName, and defaults to PreemptLowerPriority if value is unset in PriorityClass. This field is immutable. This field is available only when the PodGroupPreemptionPolicy feature gate is enabled.
	// Default: PreemptLowerPriority if unset. When Priority Admission Controller is enabled, it populates this field from PriorityClassName, and defaults to PreemptLowerPriority if value is unset in PriorityClass. This field is immutable. This field is available only when the PodGroupPreemptionPolicy feature gate is enabled.
	//
	PreemptionPolicy *string `field:"optional" json:"preemptionPolicy" yaml:"preemptionPolicy"`
	// priority is the value of priority of this pod group.
	//
	// Various system components use this field to find the priority of the pod group. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. This field is immutable.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// priorityClassName defines the priority that should be considered when scheduling this pod group.
	//
	// Controllers are expected to fill this field by copying it from a PodGroupTemplate. Otherwise, it is validated and resolved similarly to the PriorityClassName on PodGroupTemplate (i.e. if no priority class is specified, admission control can set this to the global default priority class if it exists. Otherwise, the pod group's priority will be zero). This field is immutable.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// resourceClaims defines which ResourceClaims may be shared among Pods in the group.
	//
	// Pods consume the devices allocated to a PodGroup's claim by defining a claim in its own Spec.ResourceClaims that matches the PodGroup's claim exactly. The claim must have the same name and refer to the same ResourceClaim or ResourceClaimTemplate.
	//
	// This is a beta-level field and requires that the DRAWorkloadResourceClaims feature gate is enabled.
	//
	// This field is immutable.
	ResourceClaims *[]*PodGroupResourceClaimV1Alpha3 `field:"optional" json:"resourceClaims" yaml:"resourceClaims"`
	// schedulingConstraints defines optional scheduling constraints (e.g. topology) for this PodGroup. Controllers are expected to fill this field by copying it from a PodGroupTemplate. This field is immutable. This field is only available when the TopologyAwareWorkloadScheduling feature gate is enabled.
	SchedulingConstraints *PodGroupSchedulingConstraintsV1Alpha3 `field:"optional" json:"schedulingConstraints" yaml:"schedulingConstraints"`
	// workloadRef references an optional PodGroup template within the Workload object that was used to create the PodGroup.
	//
	// This field is immutable.
	WorkloadRef *WorkloadReferenceV1Alpha3 `field:"optional" json:"workloadRef" yaml:"workloadRef"`
}

