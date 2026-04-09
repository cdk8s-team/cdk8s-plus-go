package k8s


// EndpointConditions represents the current condition of an endpoint.
type EndpointConditions struct {
	// ready indicates that this endpoint is ready to receive traffic, according to whatever system is managing the endpoint.
	//
	// A nil value should be interpreted as "true". In general, an endpoint should be marked ready if it is serving and not terminating, though this can be overridden in some cases, such as when the associated Service has set the publishNotReadyAddresses flag.
	Ready *bool `field:"optional" json:"ready" yaml:"ready"`
	// serving indicates that this endpoint is able to receive traffic, according to whatever system is managing the endpoint.
	//
	// For endpoints backed by pods, the EndpointSlice controller will mark the endpoint as serving if the pod's Ready condition is True. A nil value should be interpreted as "true".
	Serving *bool `field:"optional" json:"serving" yaml:"serving"`
	// terminating indicates that this endpoint is terminating.
	//
	// A nil value should be interpreted as "false".
	Terminating *bool `field:"optional" json:"terminating" yaml:"terminating"`
}

