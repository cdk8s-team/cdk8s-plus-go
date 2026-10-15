package k8s


// PodGroupSchedulingConstraints defines scheduling constraints (e.g. topology) for a PodGroup.
type PodGroupSchedulingConstraintsV1Alpha2 struct {
	// Topology defines the topology constraints for the pod group.
	//
	// Currently only a single topology constraint can be specified. This may change in the future.
	Topology *[]*TopologyConstraintV1Alpha2 `field:"optional" json:"topology" yaml:"topology"`
}

