package k8s


// WorkloadPodGroupSchedulingConstraints defines leaf-level scheduling constraints, such as topology.
type WorkloadPodGroupSchedulingConstraintsV1Alpha3 struct {
	// topology specifies desired topological placements for all pods within the pod group.
	//
	// If unset, no topology placement is requested.
	Topology *[]*TopologyConstraintV1Alpha3 `field:"optional" json:"topology" yaml:"topology"`
}

