package k8s


// WorkloadSpec defines the desired state of a Workload.
type WorkloadSpecV1Beta1 struct {
	// compositePodGroupTemplates is the list of CompositePodGroup templates that make up the Workload.
	//
	// The maximum number of templates is 8. This field is immutable. Exactly one of CompositePodGroupTemplates and PodGroupTemplates must be set.
	//
	// This field is used only when the CompositePodGroup feature gate is enabled.
	CompositePodGroupTemplates *[]*CompositePodGroupTemplateV1Beta1 `field:"optional" json:"compositePodGroupTemplates" yaml:"compositePodGroupTemplates"`
	// controllerRef is an optional reference to the controlling object, such as a Deployment or Job.
	//
	// This field is intended for use by tools like CLIs to provide a link back to the original workload definition. This field is immutable.
	ControllerRef *TypedLocalObjectReferenceV1Beta1 `field:"optional" json:"controllerRef" yaml:"controllerRef"`
	// podGroupTemplates is the list of templates that make up the Workload.
	//
	// The maximum number of templates is 8. Templates cannot be added or removed after the workload is created. Existing templates may still be updated where their individual fields allow it. Exactly one of CompositePodGroupTemplates and PodGroupTemplates must be set.
	PodGroupTemplates *[]*PodGroupTemplateV1Beta1 `field:"optional" json:"podGroupTemplates" yaml:"podGroupTemplates"`
}

