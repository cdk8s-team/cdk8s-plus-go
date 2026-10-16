package k8s


// WorkloadPodGroupDisruptionMode defines how individual pods within a group can be disrupted.
//
// Exactly one mode must be set.
type WorkloadPodGroupDisruptionModeV1Alpha3 struct {
	// all specifies that all pods in the group must be disrupted together.
	All interface{} `field:"optional" json:"all" yaml:"all"`
	// single specifies that pods can be disrupted independently from each other.
	Single interface{} `field:"optional" json:"single" yaml:"single"`
}

