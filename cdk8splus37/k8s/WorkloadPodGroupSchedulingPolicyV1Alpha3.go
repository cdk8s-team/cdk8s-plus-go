package k8s


// WorkloadPodGroupSchedulingPolicy defines the scheduling policy for a group of pods managed by a workload controller.
//
// Exactly one policy must be set.
type WorkloadPodGroupSchedulingPolicyV1Alpha3 struct {
	// basic specifies that standard, pod-by-pod Kubernetes scheduling behavior should be used.
	Basic interface{} `field:"optional" json:"basic" yaml:"basic"`
	// gang specifies all-or-nothing scheduling semantics.
	Gang *WorkloadPodGroupGangSchedulingPolicyV1Alpha3 `field:"optional" json:"gang" yaml:"gang"`
}

