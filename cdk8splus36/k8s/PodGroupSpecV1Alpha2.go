package k8s


// PodGroupSpec defines the desired state of a PodGroup.
type PodGroupSpecV1Alpha2 struct {
	// SchedulingPolicy defines the scheduling policy for this instance of the PodGroup.
	//
	// Controllers are expected to fill this field by copying it from a PodGroupTemplate. This field is immutable.
	SchedulingPolicy *PodGroupSchedulingPolicyV1Alpha2 `field:"required" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// DisruptionMode defines the mode in which a given PodGroup can be disrupted.
	//
	// Controllers are expected to fill this field by copying it from a PodGroupTemplate. One of Pod, PodGroup. Defaults to Pod if unset. This field is immutable. This field is available only when the WorkloadAwarePreemption feature gate is enabled.
	// Default: Pod if unset. This field is immutable. This field is available only when the WorkloadAwarePreemption feature gate is enabled.
	//
	DisruptionMode *string `field:"optional" json:"disruptionMode" yaml:"disruptionMode"`
	// PodGroupTemplateRef references an optional PodGroup template within other object (e.g. Workload) that was used to create the PodGroup. This field is immutable.
	PodGroupTemplateRef *PodGroupTemplateReferenceV1Alpha2 `field:"optional" json:"podGroupTemplateRef" yaml:"podGroupTemplateRef"`
	// Priority is the value of priority of this pod group.
	//
	// Various system components use this field to find the priority of the pod group. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. This field is immutable. This field is available only when the WorkloadAwarePreemption feature gate is enabled.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// PriorityClassName defines the priority that should be considered when scheduling this pod group.
	//
	// Controllers are expected to fill this field by copying it from a PodGroupTemplate. Otherwise, it is validated and resolved similarly to the PriorityClassName on PodGroupTemplate (i.e. if no priority class is specified, admission control can set this to the global default priority class if it exists. Otherwise, the pod group's priority will be zero). This field is immutable. This field is available only when the WorkloadAwarePreemption feature gate is enabled.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// ResourceClaims defines which ResourceClaims may be shared among Pods in the group.
	//
	// Pods consume the devices allocated to a PodGroup's claim by defining a claim in its own Spec.ResourceClaims that matches the PodGroup's claim exactly. The claim must have the same name and refer to the same ResourceClaim or ResourceClaimTemplate.
	//
	// This is an alpha-level field and requires that the DRAWorkloadResourceClaims feature gate is enabled.
	//
	// This field is immutable.
	ResourceClaims *[]*PodGroupResourceClaimV1Alpha2 `field:"optional" json:"resourceClaims" yaml:"resourceClaims"`
	// SchedulingConstraints defines optional scheduling constraints (e.g. topology) for this PodGroup. Controllers are expected to fill this field by copying it from a PodGroupTemplate. This field is immutable. This field is only available when the TopologyAwareWorkloadScheduling feature gate is enabled.
	SchedulingConstraints *PodGroupSchedulingConstraintsV1Alpha2 `field:"optional" json:"schedulingConstraints" yaml:"schedulingConstraints"`
}

