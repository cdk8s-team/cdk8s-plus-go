// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


type FromServiceAccountNameOptions struct {
	// The name of the namespace the service account belongs to.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
}

