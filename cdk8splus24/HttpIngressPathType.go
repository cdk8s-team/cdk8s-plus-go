// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


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

