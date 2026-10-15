package k8s


// PodGroupTemplate represents a template for a set of pods with a scheduling policy.
type PodGroupTemplateV1Alpha2 struct {
	// Name is a unique identifier for the PodGroupTemplate within the Workload.
	//
	// It must be a DNS label. This field is immutable.
	Name *string `field:"required" json:"name" yaml:"name"`
	// SchedulingPolicy defines the scheduling policy for this PodGroupTemplate.
	SchedulingPolicy *PodGroupSchedulingPolicyV1Alpha2 `field:"required" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// DisruptionMode defines the mode in which a given PodGroup can be disrupted.
	//
	// One of Pod, PodGroup. This field is available only when the WorkloadAwarePreemption feature gate is enabled.
	DisruptionMode *string `field:"optional" json:"disruptionMode" yaml:"disruptionMode"`
	// Priority is the value of priority of pod groups created from this template.
	//
	// Various system components use this field to find the priority of the pod group. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. This field is available only when the WorkloadAwarePreemption feature gate is enabled.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// PriorityClassName indicates the priority that should be considered when scheduling a pod group created from this template.
	//
	// If no priority class is specified, admission control can set this to the global default priority class if it exists. Otherwise, pod groups created from this template will have the priority set to zero. This field is available only when the WorkloadAwarePreemption feature gate is enabled.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// ResourceClaims defines which ResourceClaims may be shared among Pods in the group.
	//
	// Pods consume the devices allocated to a PodGroup's claim by defining a claim in its own Spec.ResourceClaims that matches the PodGroup's claim exactly. The claim must have the same name and refer to the same ResourceClaim or ResourceClaimTemplate.
	//
	// This is an alpha-level field and requires that the DRAWorkloadResourceClaims feature gate is enabled.
	//
	// This field is immutable.
	ResourceClaims *[]*PodGroupResourceClaimV1Alpha2 `field:"optional" json:"resourceClaims" yaml:"resourceClaims"`
	// SchedulingConstraints defines optional scheduling constraints (e.g. topology) for this PodGroupTemplate. This field is only available when the TopologyAwareWorkloadScheduling feature gate is enabled.
	SchedulingConstraints *PodGroupSchedulingConstraintsV1Alpha2 `field:"optional" json:"schedulingConstraints" yaml:"schedulingConstraints"`
}

