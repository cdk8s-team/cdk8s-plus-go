// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Options for exposing a service using an ingress.
type ExposeServiceViaIngressOptions struct {
	// The ingress to add rules to.
	Ingress Ingress `field:"optional" json:"ingress" yaml:"ingress"`
	// The type of the path.
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

