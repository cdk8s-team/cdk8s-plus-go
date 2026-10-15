package k8s


// PodGroupResourceClaim references exactly one ResourceClaim, either directly or by naming a ResourceClaimTemplate which is then turned into a ResourceClaim for the PodGroup.
//
// It adds a name to it that uniquely identifies the ResourceClaim inside the PodGroup. Pods that need access to the ResourceClaim define a matching reference in its own Spec.ResourceClaims. The Pod's claim must match all fields of the PodGroup's claim exactly.
type PodGroupResourceClaimV1Alpha2 struct {
	// Name uniquely identifies this resource claim inside the PodGroup.
	//
	// This must be a DNS_LABEL.
	Name *string `field:"required" json:"name" yaml:"name"`
	// ResourceClaimName is the name of a ResourceClaim object in the same namespace as this PodGroup.
	//
	// The ResourceClaim will be reserved for the PodGroup instead of its individual pods.
	//
	// Exactly one of ResourceClaimName and ResourceClaimTemplateName must be set.
	ResourceClaimName *string `field:"optional" json:"resourceClaimName" yaml:"resourceClaimName"`
	// ResourceClaimTemplateName is the name of a ResourceClaimTemplate object in the same namespace as this PodGroup.
	//
	// The template will be used to create a new ResourceClaim, which will be bound to this PodGroup. When this PodGroup is deleted, the ResourceClaim will also be deleted. The PodGroup name and resource name, along with a generated component, will be used to form a unique name for the ResourceClaim, which will be recorded in podgroup.status.resourceClaimStatuses.
	//
	// This field is immutable and no changes will be made to the corresponding ResourceClaim by the control plane after creating the ResourceClaim.
	//
	// Exactly one of ResourceClaimName and ResourceClaimTemplateName must be set.
	ResourceClaimTemplateName *string `field:"optional" json:"resourceClaimTemplateName" yaml:"resourceClaimTemplateName"`
}

