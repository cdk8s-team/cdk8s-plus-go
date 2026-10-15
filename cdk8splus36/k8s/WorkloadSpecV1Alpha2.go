package k8s


// WorkloadSpec defines the desired state of a Workload.
type WorkloadSpecV1Alpha2 struct {
	// PodGroupTemplates is the list of templates that make up the Workload.
	//
	// The maximum number of templates is 8. This field is immutable.
	PodGroupTemplates *[]*PodGroupTemplateV1Alpha2 `field:"required" json:"podGroupTemplates" yaml:"podGroupTemplates"`
	// ControllerRef is an optional reference to the controlling object, such as a Deployment or Job.
	//
	// This field is intended for use by tools like CLIs to provide a link back to the original workload definition. This field is immutable.
	ControllerRef *TypedLocalObjectReferenceV1Alpha2 `field:"optional" json:"controllerRef" yaml:"controllerRef"`
}

