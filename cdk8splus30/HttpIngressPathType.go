package cdk8splus30


// Specify how the path is matched against request paths.
// See: https://kubernetes.io/docs/concepts/services-networking/ingress/#path-types
//
type HttpIngressPathType string

const (
	// Matches the URL path exactly.
	HttpIngressPathType_PREFIX HttpIngressPathType = "PREFIX"
	// Matches based on a URL path prefix split by '/'.
	HttpIngressPathType_EXACT HttpIngressPathType = "EXACT"
	// Matching is specified by the underlying IngressClass.
	HttpIngressPathType_IMPLEMENTATION_SPECIFIC HttpIngressPathType = "IMPLEMENTATION_SPECIFIC"
)

