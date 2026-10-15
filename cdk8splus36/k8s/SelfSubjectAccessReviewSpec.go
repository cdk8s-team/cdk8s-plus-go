package k8s


// SelfSubjectAccessReviewSpec is a description of the access request.
//
// Exactly one of resourceAttributes and nonResourceAttributes must be set.
type SelfSubjectAccessReviewSpec struct {
	// nonResourceAttributes describes information for a non-resource access request.
	NonResourceAttributes *NonResourceAttributes `field:"optional" json:"nonResourceAttributes" yaml:"nonResourceAttributes"`
	// resourceAttributes describes information for a resource access request.
	ResourceAttributes *ResourceAttributes `field:"optional" json:"resourceAttributes" yaml:"resourceAttributes"`
}

