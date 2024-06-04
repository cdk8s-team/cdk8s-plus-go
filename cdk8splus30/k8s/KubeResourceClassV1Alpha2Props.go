package k8s


// ResourceClass is used by administrators to influence how resources are allocated.
//
// This is an alpha type and requires enabling the DynamicResourceAllocation feature gate.
type KubeResourceClassV1Alpha2Props struct {
	// DriverName defines the name of the dynamic resource driver that is used for allocation of a ResourceClaim that uses this class.
	//
	// Resource drivers have a unique name in forward domain order (acme.example.com).
	DriverName *string `field:"required" json:"driverName" yaml:"driverName"`
	// Standard object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// ParametersRef references an arbitrary separate object that may hold parameters that will be used by the driver when allocating a resource that uses this class.
	//
	// A dynamic resource driver can distinguish between parameters stored here and and those stored in ResourceClaimSpec.
	ParametersRef *ResourceClassParametersReferenceV1Alpha2 `field:"optional" json:"parametersRef" yaml:"parametersRef"`
	// If and only if allocation of claims using this class is handled via structured parameters, then StructuredParameters must be set to true.
	StructuredParameters *bool `field:"optional" json:"structuredParameters" yaml:"structuredParameters"`
	// Only nodes matching the selector will be considered by the scheduler when trying to find a Node that fits a Pod when that Pod uses a ResourceClaim that has not been allocated yet.
	//
	// Setting this field is optional. If null, all nodes are candidates.
	SuitableNodes *NodeSelector `field:"optional" json:"suitableNodes" yaml:"suitableNodes"`
}

