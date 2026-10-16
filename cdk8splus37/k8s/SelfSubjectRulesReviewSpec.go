package k8s


// SelfSubjectRulesReviewSpec defines the specification for SelfSubjectRulesReview.
type SelfSubjectRulesReviewSpec struct {
	// namespace to evaluate rules for.
	//
	// Required.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

