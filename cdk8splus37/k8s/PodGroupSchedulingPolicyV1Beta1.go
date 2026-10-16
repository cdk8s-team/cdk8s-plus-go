package k8s


// PodGroupSchedulingPolicy defines the scheduling configuration for a PodGroup.
//
// Exactly one policy must be set. The policy is chosen at creation time by setting either the Basic or Gang field. The PodGroup may not change policy after creation. Fields within chosen policy may be updated after creation when their individual fields allow it.
type PodGroupSchedulingPolicyV1Beta1 struct {
	// basic specifies that the pods in this group should be scheduled using standard Kubernetes scheduling behavior.
	//
	// Setting this field at group creation time opts this group to basic scheduling; this field cannot be changed afterward.
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
	// gang specifies that the pods in this group should be scheduled using all-or-nothing semantics.
	//
	// Setting this field at group creation time opts this group to gang scheduling; this field cannot be set or unset afterward. The minCount field within Gang scheduling policy remains mutable after group creation.
	Gang *GangSchedulingPolicyV1Beta1 `field:"optional" json:"gang" yaml:"gang"`
}

