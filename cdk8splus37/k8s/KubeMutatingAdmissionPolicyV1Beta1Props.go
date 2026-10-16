package k8s


// MutatingAdmissionPolicy describes the definition of an admission mutation policy that mutates the object coming into admission chain.
type KubeMutatingAdmissionPolicyV1Beta1Props struct {
	// metadata is the standard object metadata;
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// spec defines the desired behavior of the MutatingAdmissionPolicy.
	Spec *MutatingAdmissionPolicySpecV1Beta1 `field:"optional" json:"spec" yaml:"spec"`
}

