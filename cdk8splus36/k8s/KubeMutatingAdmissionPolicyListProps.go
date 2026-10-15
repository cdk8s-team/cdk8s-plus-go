package k8s


// MutatingAdmissionPolicyList is a list of MutatingAdmissionPolicy.
type KubeMutatingAdmissionPolicyListProps struct {
	// List of ValidatingAdmissionPolicy.
	Items *[]*KubeMutatingAdmissionPolicyProps `field:"required" json:"items" yaml:"items"`
	// metadata is the standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

