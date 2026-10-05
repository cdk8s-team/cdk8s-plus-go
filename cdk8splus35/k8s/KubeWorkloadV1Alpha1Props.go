package k8s


// Workload allows for expressing scheduling constraints that should be used when managing lifecycle of workloads from scheduling perspective, including scheduling, preemption, eviction and other phases.
type KubeWorkloadV1Alpha1Props struct {
	// Spec defines the desired behavior of a Workload.
	Spec *WorkloadSpecV1Alpha1 `field:"required" json:"spec" yaml:"spec"`
	// Standard object's metadata.
	//
	// Name must be a DNS subdomain.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

