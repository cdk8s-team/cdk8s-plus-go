package k8s


// PodDisruptionBudget is an object to define the max disruption that can be caused to a collection of pods.
type KubePodDisruptionBudgetV1Beta1Props struct {
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Specification of the desired behavior of the PodDisruptionBudget.
	Spec *PodDisruptionBudgetSpecV1Beta1 `field:"optional" json:"spec" yaml:"spec"`
}

