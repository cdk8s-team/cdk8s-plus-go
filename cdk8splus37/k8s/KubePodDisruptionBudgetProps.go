package k8s


// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods.
type KubePodDisruptionBudgetProps struct {
	// metadata is the standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// spec is the specification of the desired behavior of the PodDisruptionBudget.
	Spec *PodDisruptionBudgetSpec `field:"optional" json:"spec" yaml:"spec"`
}

