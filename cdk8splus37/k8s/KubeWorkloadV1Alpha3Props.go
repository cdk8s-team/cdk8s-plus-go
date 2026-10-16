package k8s


// Workload allows for expressing scheduling constraints that should be used when managing the lifecycle of workloads from the scheduling perspective, including scheduling, preemption, eviction and other phases.
//
// Workload API enablement is toggled by the GenericWorkload feature gate.
type KubeWorkloadV1Alpha3Props struct {
	// spec defines the desired behavior of a Workload.
	Spec *WorkloadSpecV1Alpha3 `field:"required" json:"spec" yaml:"spec"`
	// metadata is the standard object metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

