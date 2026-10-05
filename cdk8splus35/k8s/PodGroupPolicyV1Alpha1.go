package k8s


// PodGroupPolicy defines the scheduling configuration for a PodGroup.
type PodGroupPolicyV1Alpha1 struct {
	// Basic specifies that the pods in this group should be scheduled using standard Kubernetes scheduling behavior.
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
	// Gang specifies that the pods in this group should be scheduled using all-or-nothing semantics.
	Gang *GangSchedulingPolicyV1Alpha1 `field:"optional" json:"gang" yaml:"gang"`
}

