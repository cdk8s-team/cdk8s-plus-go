package k8s


// PodGroupTemplateReference references a PodGroup template defined in some object (e.g. Workload). Exactly one reference must be set.
type PodGroupTemplateReferenceV1Alpha2 struct {
	// Workload references the PodGroupTemplate within the Workload object that was used to create the PodGroup.
	Workload *WorkloadPodGroupTemplateReferenceV1Alpha2 `field:"optional" json:"workload" yaml:"workload"`
}

