// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


type FromServiceAccountNameOptions struct {
	// The name of the namespace the service account belongs to.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
}

