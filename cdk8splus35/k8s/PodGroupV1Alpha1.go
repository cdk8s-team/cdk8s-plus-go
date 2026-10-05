package k8s


// PodGroup represents a set of pods with a common scheduling policy.
type PodGroupV1Alpha1 struct {
	// Name is a unique identifier for the PodGroup within the Workload.
	//
	// It must be a DNS label. This field is immutable.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Policy defines the scheduling policy for this PodGroup.
	Policy *PodGroupPolicyV1Alpha1 `field:"required" json:"policy" yaml:"policy"`
}

