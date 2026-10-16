package k8s


// MutatingAdmissionPolicy describes the definition of an admission mutation policy that mutates the object coming into admission chain.
type KubeMutatingAdmissionPolicyProps struct {
	// metadata is the standard object metadata;
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// spec defines the desired behavior of the MutatingAdmissionPolicy.
	Spec *MutatingAdmissionPolicySpec `field:"optional" json:"spec" yaml:"spec"`
}

