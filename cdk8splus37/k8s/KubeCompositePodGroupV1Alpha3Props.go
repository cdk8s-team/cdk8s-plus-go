package k8s


// CompositePodGroup represents a runtime instance of pod groups grouped together.
//
// CompositePodGroups are created by workload controllers (LWS, JobSet, etc...) from Workload.compositePodGroupTemplates. CompositePodGroup API enablement is toggled by the CompositePodGroup feature gate.
type KubeCompositePodGroupV1Alpha3Props struct {
	// spec defines the desired state of the CompositePodGroup.
	Spec *CompositePodGroupSpecV1Alpha3 `field:"required" json:"spec" yaml:"spec"`
	// metadata is the standard object metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

