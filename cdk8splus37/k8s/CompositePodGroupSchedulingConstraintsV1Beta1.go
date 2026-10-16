package k8s


// CompositePodGroupSchedulingConstraints defines scheduling constraints (e.g. topology) for a CompositePodGroup.
type CompositePodGroupSchedulingConstraintsV1Beta1 struct {
	// topology defines the topology constraints for the composite pod group.
	//
	// Currently only a single topology constraint can be specified. This may change in the future.
	Topology *[]*TopologyConstraintV1Beta1 `field:"optional" json:"topology" yaml:"topology"`
}

