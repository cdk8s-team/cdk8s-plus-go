package k8s


// PriorityLevelConfiguration represents the configuration of a priority level.
type KubePriorityLevelConfigurationProps struct {
	// `spec` is the specification of the desired behavior of a "request-priority".
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	Spec *PriorityLevelConfigurationSpec `field:"required" json:"spec" yaml:"spec"`
	// `metadata` is the standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

