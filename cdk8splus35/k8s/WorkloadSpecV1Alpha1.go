package k8s


// WorkloadSpec defines the desired state of a Workload.
type WorkloadSpecV1Alpha1 struct {
	// PodGroups is the list of pod groups that make up the Workload.
	//
	// The maximum number of pod groups is 8. This field is immutable.
	PodGroups *[]*PodGroupV1Alpha1 `field:"required" json:"podGroups" yaml:"podGroups"`
	// ControllerRef is an optional reference to the controlling object, such as a Deployment or Job.
	//
	// This field is intended for use by tools like CLIs to provide a link back to the original workload definition. When set, it cannot be changed.
	ControllerRef *TypedLocalObjectReferenceV1Alpha1 `field:"optional" json:"controllerRef" yaml:"controllerRef"`
}

