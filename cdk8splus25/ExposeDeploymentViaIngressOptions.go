package cdk8splus25


// Options for exposing a deployment via an ingress.
type ExposeDeploymentViaIngressOptions struct {
	// The name of the service to expose.
	//
	// If you'd like to expose the deployment multiple times,
	// you must explicitly set a name starting from the second expose call.
	// Default: - auto generated.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ports that the service should bind to.
	// Default: - extracted from the deployment.
	//
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// The type of the exposed service.
	// Default: - ClusterIP.
	//
	ServiceType ServiceType `field:"optional" json:"serviceType" yaml:"serviceType"`
	// The ingress to add rules to.
	// Default: - An ingress will be automatically created.
	//
	Ingress Ingress `field:"optional" json:"ingress" yaml:"ingress"`
	// The type of the path.
	// Default: HttpIngressPathType.PREFIX
	//
	PathType HttpIngressPathType `field:"optional" json:"pathType" yaml:"pathType"`
}

