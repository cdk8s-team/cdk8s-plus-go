// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


type ConnectionScheme string

const (
	// Use HTTP request for connecting to host.
	ConnectionScheme_HTTP ConnectionScheme = "HTTP"
	// Use HTTPS request for connecting to host.
	ConnectionScheme_HTTPS ConnectionScheme = "HTTPS"
)

