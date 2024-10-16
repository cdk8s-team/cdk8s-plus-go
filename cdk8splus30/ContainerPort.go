package cdk8splus30


// Represents a network port in a single container.
type ContainerPort struct {
	// Number of port to expose on the pod's IP address.
	//
	// This must be a valid port number, 0 < x < 65536.
	Number *float64 `field:"required" json:"number" yaml:"number"`
	// What host IP to bind the external port to.
	// Default: - 127.0.0.1.
	//
	HostIp *string `field:"optional" json:"hostIp" yaml:"hostIp"`
	// Number of port to expose on the host.
	//
	// If specified, this must be a valid port number, 0 < x < 65536.
	// Most containers do not need this.
	// Default: - auto generated by kubernetes and might change on restarts.
	//
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// If specified, this must be an IANA_SVC_NAME and unique within the pod.
	//
	// Each named port in a pod must have a unique name.
	// Name for the port that can be referred to by services.
	// Default: - port is not named.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Protocol for port.
	//
	// Must be UDP, TCP, or SCTP. Defaults to "TCP".
	// Default: Protocol.TCP
	//
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

