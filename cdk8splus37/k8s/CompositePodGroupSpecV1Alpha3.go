package k8s


// CompositePodGroupSpec defines the desired state of CompositePodGroup.
type CompositePodGroupSpecV1Alpha3 struct {
	// schedulingPolicy defines the scheduling policy for this instance of the CompositePodGroup.
	//
	// Controllers are expected to fill this field by copying it from a CompositePodGroupTemplate. This field is immutable.
	SchedulingPolicy *CompositePodGroupSchedulingPolicyV1Alpha3 `field:"required" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// workloadRef references an optional CompositePodGroup template within the Workload object that was used to create the CompositePodGroup.
	//
	// This field is required. This field is immutable.
	WorkloadRef *WorkloadReferenceV1Alpha3 `field:"required" json:"workloadRef" yaml:"workloadRef"`
	// disruptionMode defines the mode in which a given CompositePodGroup can be disrupted.
	//
	// Controllers are expected to fill this field by copying it from a CompositePodGroupTemplate. One of Single, All. Defaults to Single if unset. This field is immutable.
	// Default: Single if unset. This field is immutable.
	//
	DisruptionMode *CompositeDisruptionModeV1Alpha3 `field:"optional" json:"disruptionMode" yaml:"disruptionMode"`
	// parentCompositePodGroupName contains the name of the parent composite pod group within the same namespace as this composite pod group.
	//
	// It must be a DNS name. If it's nil, then this composite pod group is a root of a workload's hierarchy. This field is immutable.
	ParentCompositePodGroupName *string `field:"optional" json:"parentCompositePodGroupName" yaml:"parentCompositePodGroupName"`
	// preemptionPolicy is the Policy for preempting pods/podgroups with lower priority.
	//
	// One of Never, PreemptLowerPriority. Defaults to PreemptLowerPriority if unset. When Priority Admission Controller is enabled, it populates this field from PriorityClassName, and defaults to PreemptLowerPriority if value is unset in PriorityClass. This field is immutable. This field is available only when the PodGroupPreemptionPolicy feature gate is enabled.
	// Default: PreemptLowerPriority if unset. When Priority Admission Controller is enabled, it populates this field from PriorityClassName, and defaults to PreemptLowerPriority if value is unset in PriorityClass. This field is immutable. This field is available only when the PodGroupPreemptionPolicy feature gate is enabled.
	//
	PreemptionPolicy *string `field:"optional" json:"preemptionPolicy" yaml:"preemptionPolicy"`
	// priority is the value of priority of this composite pod group.
	//
	// Various system components use this field to find the priority of the composite pod group. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. This field is immutable.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// priorityClassName defines the priority that should be considered when scheduling this CompositePodGroup.
	//
	// Controllers are expected to fill this field by copying it from a CompositePodGroupTemplate. If left unspecified, it is validated and resolved similarly to the PriorityClassName field in Pods (i.e. if no priority class is specified, admission control can set this to the global default priority class if it exists. Otherwise, the composite pod group's priority will be zero). This field is immutable.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// schedulingConstraints defines optional scheduling constraints (e.g. topology) for this CompositePodGroup. Controllers are expected to fill this field by copying it from a CompositePodGroupTemplate. This field is immutable.
	SchedulingConstraints *CompositePodGroupSchedulingConstraintsV1Alpha3 `field:"optional" json:"schedulingConstraints" yaml:"schedulingConstraints"`
}

