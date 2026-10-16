package k8s


// WorkloadPodGroupResourceClaim references a dynamic resource claim that is shared across pods in the group.
type WorkloadPodGroupResourceClaimV1Alpha3 struct {
	// name uniquely identifies this resource claim inside the group.
	//
	// This field is required. It must be a DNS_LABEL.
	Name *string `field:"required" json:"name" yaml:"name"`
	// resourceClaimName is the name of a ResourceClaim object in the same namespace.
	//
	// This field is optional. If it is not specified, no resource claim is used. If set, it must be a DNS subdomain.
	ResourceClaimName *string `field:"optional" json:"resourceClaimName" yaml:"resourceClaimName"`
	// resourceClaimTemplateName is the name of a ResourceClaimTemplate object in the same namespace.
	//
	// This field is optional. If it is not specified, no resource claim template is used. If set, it must be a DNS subdomain.
	ResourceClaimTemplateName *string `field:"optional" json:"resourceClaimTemplateName" yaml:"resourceClaimTemplateName"`
}

