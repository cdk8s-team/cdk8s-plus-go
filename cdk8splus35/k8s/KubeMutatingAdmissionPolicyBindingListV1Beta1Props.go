package k8s


// MutatingAdmissionPolicyBindingList is a list of MutatingAdmissionPolicyBinding.
type KubeMutatingAdmissionPolicyBindingListV1Beta1Props struct {
	// List of PolicyBinding.
	Items *[]*KubeMutatingAdmissionPolicyBindingV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

