// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Options for exposing a deployment via an ingress.
type ExposeDeploymentViaIngressOptions struct {
	// The name of the service to expose.
	//
	// If you'd like to expose the deployment multiple times,
	// you must explicitly set a name starting from the second expose call.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ports that the service should bind to.
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// The type of the exposed service.
	ServiceType ServiceType `field:"optional" json:"serviceType" yaml:"serviceType"`
	// The ingress to add rules to.
	Ingress Ingress `field:"optional" json:"ingress" yaml:"ingress"`
	// The type of the path.
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

