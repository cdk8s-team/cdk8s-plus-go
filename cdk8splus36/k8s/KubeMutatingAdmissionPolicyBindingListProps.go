package k8s


// MutatingAdmissionPolicyBindingList is a list of MutatingAdmissionPolicyBinding.
type KubeMutatingAdmissionPolicyBindingListProps struct {
	// List of PolicyBinding.
	Items *[]*KubeMutatingAdmissionPolicyBindingProps `field:"required" json:"items" yaml:"items"`
	// metadata is the standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

