package k8s


// WorkloadPodGroupTemplateReference references the PodGroupTemplate within the Workload object.
type WorkloadPodGroupTemplateReferenceV1Alpha2 struct {
	// PodGroupTemplateName defines the PodGroupTemplate name within the Workload object.
	PodGroupTemplateName *string `field:"required" json:"podGroupTemplateName" yaml:"podGroupTemplateName"`
	// WorkloadName defines the name of the Workload object.
	WorkloadName *string `field:"required" json:"workloadName" yaml:"workloadName"`
}

