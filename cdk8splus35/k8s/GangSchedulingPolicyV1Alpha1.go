package k8s


// GangSchedulingPolicy defines the parameters for gang scheduling.
type GangSchedulingPolicyV1Alpha1 struct {
	// MinCount is the minimum number of pods that must be schedulable or scheduled at the same time for the scheduler to admit the entire group.
	//
	// It must be a positive integer.
	MinCount *float64 `field:"required" json:"minCount" yaml:"minCount"`
}

