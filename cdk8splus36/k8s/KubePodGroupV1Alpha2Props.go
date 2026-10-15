package k8s


// PodGroup represents a runtime instance of pods grouped together.
//
// PodGroups are created by workload controllers (Job, LWS, JobSet, etc...) from Workload.podGroupTemplates. PodGroup API enablement is toggled by the GenericWorkload feature gate.
type KubePodGroupV1Alpha2Props struct {
	// Spec defines the desired state of the PodGroup.
	Spec *PodGroupSpecV1Alpha2 `field:"required" json:"spec" yaml:"spec"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

