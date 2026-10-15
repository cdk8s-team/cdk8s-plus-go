package k8s


// Workload allows for expressing scheduling constraints that should be used when managing the lifecycle of workloads from the scheduling perspective, including scheduling, preemption, eviction and other phases.
//
// Workload API enablement is toggled by the GenericWorkload feature gate.
type KubeWorkloadV1Alpha2Props struct {
	// Spec defines the desired behavior of a Workload.
	Spec *WorkloadSpecV1Alpha2 `field:"required" json:"spec" yaml:"spec"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

