package k8s


// TokenReview attempts to authenticate a token to a known user.
//
// Note: TokenReview requests may be cached by the webhook token authenticator plugin in the kube-apiserver.
type KubeTokenReviewProps struct {
	// spec holds information about the request being evaluated.
	Spec *TokenReviewSpec `field:"required" json:"spec" yaml:"spec"`
	// metadata is the standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

