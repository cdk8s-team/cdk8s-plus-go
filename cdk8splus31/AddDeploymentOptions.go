package cdk8splus31


// Options to add a deployment to a service.
type AddDeploymentOptions struct {
	// The name of this port within the service.
	//
	// This must be a DNS_LABEL. All
	// ports within a ServiceSpec must have unique names. This maps to the 'Name'
	// field in EndpointPort objects. Optional if only one ServicePort is defined
	// on this service.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The port on each node on which this service is exposed when type=NodePort or LoadBalancer.
	//
	// Usually assigned by the system. If specified, it will be
	// allocated to the service if unused or else creation of the service will
	// fail. Default is to auto-allocate a port if the ServiceType of this Service
	// requires one.
	// See: https://kubernetes.io/docs/concepts/services-networking/service/#type-nodeport
	//
	// Default: - auto-allocate a port if the ServiceType of this Service requires one.
	//
	NodePort *float64 `field:"optional" json:"nodePort" yaml:"nodePort"`
	// The IP protocol for this port.
	//
	// Supports "TCP", "UDP", and "SCTP". Default is TCP.
	// Default: Protocol.TCP
	//
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
	// The port number the service will redirect to.
	// Default: - The value of `port` will be used.
	//
	TargetPort *float64 `field:"optional" json:"targetPort" yaml:"targetPort"`
	// The port number the service will bind to.
	// Default: - Copied from the first container of the deployment.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

