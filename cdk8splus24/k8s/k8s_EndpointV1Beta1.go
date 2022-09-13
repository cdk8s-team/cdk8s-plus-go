package k8s


// Endpoint represents a single logical "backend" implementing a service.
type EndpointV1Beta1 struct {
	// addresses of this endpoint.
	//
	// The contents of this field are interpreted according to the corresponding EndpointSlice addressType field. Consumers must handle different types of addresses in the context of their own capabilities. This must contain at least one address but no more than 100. These are all assumed to be fungible and clients may choose to only use the first element. Refer to: https://issue.k8s.io/106267
	Addresses *[]*string `field:"required" json:"addresses" yaml:"addresses"`
	// conditions contains information about the current status of the endpoint.
	Conditions *EndpointConditionsV1Beta1 `field:"optional" json:"conditions" yaml:"conditions"`
	// hints contains information associated with how an endpoint should be consumed.
	Hints *EndpointHintsV1Beta1 `field:"optional" json:"hints" yaml:"hints"`
	// hostname of this endpoint.
	//
	// This field may be used by consumers of endpoints to distinguish endpoints from each other (e.g. in DNS names). Multiple endpoints which use the same hostname should be considered fungible (e.g. multiple A values in DNS). Must be lowercase and pass DNS Label (RFC 1123) validation.
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// nodeName represents the name of the Node hosting this endpoint.
	//
	// This can be used to determine endpoints local to a Node. This field can be enabled with the EndpointSliceNodeName feature gate.
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
	// targetRef is a reference to a Kubernetes object that represents this endpoint.
	TargetRef *ObjectReference `field:"optional" json:"targetRef" yaml:"targetRef"`
	// topology contains arbitrary topology information associated with the endpoint.
	//
	// These key/value pairs must conform with the label format. https://kubernetes.io/docs/concepts/overview/working-with-objects/labels Topology may include a maximum of 16 key/value pairs. This includes, but is not limited to the following well known keys: * kubernetes.io/hostname: the value indicates the hostname of the node
	// where the endpoint is located. This should match the corresponding
	// node label.
	// * topology.kubernetes.io/zone: the value indicates the zone where the
	// endpoint is located. This should match the corresponding node label.
	// * topology.kubernetes.io/region: the value indicates the region where the
	// endpoint is located. This should match the corresponding node label.
	// This field is deprecated and will be removed in future api versions.
	Topology *map[string]*string `field:"optional" json:"topology" yaml:"topology"`
}

