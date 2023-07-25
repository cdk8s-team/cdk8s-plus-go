package cdk8splus27


// Options for `Deployment.exposeViaService`.
type DeploymentExposeViaServiceOptions struct {
	// The name of the service to expose.
	//
	// If you'd like to expose the deployment multiple times,
	// you must explicitly set a name starting from the second expose call.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The ports that the service should bind to.
	Ports *[]*ServicePort `field:"optional" json:"ports" yaml:"ports"`
	// The type of the exposed service.
	ServiceType ServiceType `field:"optional" json:"serviceType" yaml:"serviceType"`
}

