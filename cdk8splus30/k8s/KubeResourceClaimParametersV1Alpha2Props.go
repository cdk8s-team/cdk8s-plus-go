package k8s


// ResourceClaimParameters defines resource requests for a ResourceClaim in an in-tree format understood by Kubernetes.
type KubeResourceClaimParametersV1Alpha2Props struct {
	// DriverRequests describes all resources that are needed for the allocated claim.
	//
	// A single claim may use resources coming from different drivers. For each driver, this array has at most one entry which then may have one or more per-driver requests.
	//
	// May be empty, in which case the claim can always be allocated.
	DriverRequests *[]*DriverRequestsV1Alpha2 `field:"optional" json:"driverRequests" yaml:"driverRequests"`
	// If this object was created from some other resource, then this links back to that resource.
	//
	// This field is used to find the in-tree representation of the claim parameters when the parameter reference of the claim refers to some unknown type.
	GeneratedFrom *ResourceClaimParametersReferenceV1Alpha2 `field:"optional" json:"generatedFrom" yaml:"generatedFrom"`
	// Standard object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// Shareable indicates whether the allocated claim is meant to be shareable by multiple consumers at the same time.
	Shareable *bool `field:"optional" json:"shareable" yaml:"shareable"`
}

