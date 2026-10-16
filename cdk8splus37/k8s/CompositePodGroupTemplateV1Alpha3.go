package k8s


// CompositePodGroupTemplate represents a template for a CompositePodGroup with a scheduling policy.
type CompositePodGroupTemplateV1Alpha3 struct {
	// name is a unique identifier for the CompositePodGroupTemplate within the Workload.
	//
	// It must be a DNS label. This field is required.
	Name *string `field:"required" json:"name" yaml:"name"`
	// schedulingPolicy defines the scheduling policy for this template.
	SchedulingPolicy *CompositePodGroupSchedulingPolicyV1Alpha3 `field:"required" json:"schedulingPolicy" yaml:"schedulingPolicy"`
	// compositePodGroupTemplates is the list of templates for children CompositePodGroups.
	//
	// The maximum number of templates is 8. At least one entry in CompositePodGroupTemplates or PodGroupTemplates must be set.
	CompositePodGroupTemplates *[]*CompositePodGroupTemplateV1Alpha3 `field:"optional" json:"compositePodGroupTemplates" yaml:"compositePodGroupTemplates"`
	// disruptionMode defines the mode in which a given CompositePodGroup can be disrupted.
	//
	// One of Single, All. This field is immutable.
	DisruptionMode *CompositeDisruptionModeV1Alpha3 `field:"optional" json:"disruptionMode" yaml:"disruptionMode"`
	// podGroupTemplates is the list of templates for children PodGroups.
	//
	// The maximum number of templates is 8. At least one entry in CompositePodGroupTemplates or PodGroupTemplates must be set.
	PodGroupTemplates *[]*PodGroupTemplateV1Alpha3 `field:"optional" json:"podGroupTemplates" yaml:"podGroupTemplates"`
	// preemptionPolicy is the Policy for preempting pods/podgroups with lower priority.
	//
	// One of Never, PreemptLowerPriority. This field is immutable. This field is available only when the PodGroupPreemptionPolicy feature gate is enabled.
	PreemptionPolicy *string `field:"optional" json:"preemptionPolicy" yaml:"preemptionPolicy"`
	// priority is the value of priority of composite pod groups created from this template.
	//
	// Various system components use this field to find the priority of the composite pod group. When Priority Admission Controller is enabled, it prevents users from setting this field. The admission controller populates this field from PriorityClassName. The higher the value, the higher the priority. This field is immutable.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// priorityClassName indicates the priority that should be considered when scheduling a composite pod group created from this template.
	//
	// If no priority class is specified, admission control can set this to the global default priority class if it exists. Otherwise, composite pod groups created from this template will have the priority set to zero. This field is immutable.
	PriorityClassName *string `field:"optional" json:"priorityClassName" yaml:"priorityClassName"`
	// schedulingConstraints defines optional scheduling constraints (e.g. topology) for this CompositePodGroupTemplate. This field is immutable.
	SchedulingConstraints *CompositePodGroupSchedulingConstraintsV1Alpha3 `field:"optional" json:"schedulingConstraints" yaml:"schedulingConstraints"`
}

