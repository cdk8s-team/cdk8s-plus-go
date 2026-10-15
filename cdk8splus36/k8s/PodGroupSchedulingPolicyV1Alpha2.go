package k8s


// PodGroupSchedulingPolicy defines the scheduling configuration for a PodGroup.
//
// Exactly one policy must be set.
type PodGroupSchedulingPolicyV1Alpha2 struct {
	// Basic specifies that the pods in this group should be scheduled using standard Kubernetes scheduling behavior.
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
	// Gang specifies that the pods in this group should be scheduled using all-or-nothing semantics.
	Gang *GangSchedulingPolicyV1Alpha2 `field:"optional" json:"gang" yaml:"gang"`
}

