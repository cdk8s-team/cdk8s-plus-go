package k8s


// CompositePodGroupSchedulingPolicy defines the scheduling configuration for a CompositePodGroup.
//
// Exactly one policy must be set.
type CompositePodGroupSchedulingPolicyV1Alpha3 struct {
	// basic specifies that the groups of this composite group should be scheduled independently.
	//
	// This field is immutable.
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
	// gang specifies that the groups of this composite group should be scheduled using all-or-nothing semantics.
	Gang *CompositeGangSchedulingPolicyV1Alpha3 `field:"optional" json:"gang" yaml:"gang"`
}

