// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Policy rule of a `ClusterRole.
type ClusterRolePolicyRule struct {
	// Endpoints this rule applies to.
	//
	// Can be either api resources
	// or non api resources.
	Endpoints *[]IApiEndpoint `field:"required" json:"endpoints" yaml:"endpoints"`
	// Verbs to allow.
	//
	// (e.g ['get', 'watch'])
	Verbs *[]*string `field:"required" json:"verbs" yaml:"verbs"`
}

