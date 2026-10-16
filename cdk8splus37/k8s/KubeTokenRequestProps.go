package k8s


// TokenRequest requests a token for a given service account.
type KubeTokenRequestProps struct {
	// metadata is the standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// spec holds information about the request being evaluated.
	Spec *TokenRequestSpec `field:"optional" json:"spec" yaml:"spec"`
}

