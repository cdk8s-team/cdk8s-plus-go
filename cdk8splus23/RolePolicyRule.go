// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Policy rule of a `Role.
type RolePolicyRule struct {
	// Resources this rule applies to.
	Resources *[]IApiResource `field:"required" json:"resources" yaml:"resources"`
	// Verbs to allow.
	//
	// (e.g ['get', 'watch'])
	Verbs *[]*string `field:"required" json:"verbs" yaml:"verbs"`
}

