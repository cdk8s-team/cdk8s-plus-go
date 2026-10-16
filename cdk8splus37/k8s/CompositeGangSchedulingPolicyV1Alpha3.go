package k8s


// CompositeGangSchedulingPolicy indicates that the groups belonging to the composite group should be scheduled using all-or-nothing semantics.
type CompositeGangSchedulingPolicyV1Alpha3 struct {
	// minGroupCount is the minimum number of child groups that must be schedulable or scheduled at the same time for the scheduler to admit the entire group.
	//
	// It must be a positive integer.
	MinGroupCount *float64 `field:"required" json:"minGroupCount" yaml:"minGroupCount"`
}

