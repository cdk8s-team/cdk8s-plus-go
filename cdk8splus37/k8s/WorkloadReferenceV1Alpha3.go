package k8s


// WorkloadReference references the Workload object together with the template that was used to create a particular PodGroup.
type WorkloadReferenceV1Alpha3 struct {
	// templateName is the name of a template within the Workload object that was used to create a pod group.
	//
	// It must be a DNS label. This field is required.
	TemplateName *string `field:"required" json:"templateName" yaml:"templateName"`
	// workloadName is the name of the Workload object that contains a template that was used when creating a pod group.
	//
	// It must be a DNS name. This field is required.
	WorkloadName *string `field:"required" json:"workloadName" yaml:"workloadName"`
}

