// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


// Represents a network port in a single container.
type ContainerPort struct {
	// Number of port to expose on the pod's IP address.
	//
	// This must be a valid port number, 0 < x < 65536.
	Number *float64 `field:"required" json:"number" yaml:"number"`
	// What host IP to bind the external port to.
	HostIp *string `field:"optional" json:"hostIp" yaml:"hostIp"`
	// Number of port to expose on the host.
	//
	// If specified, this must be a valid port number, 0 < x < 65536.
	// Most containers do not need this.
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// If specified, this must be an IANA_SVC_NAME and unique within the pod.
	//
	// Each named port in a pod must have a unique name.
	// Name for the port that can be referred to by services.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Protocol for port.
	//
	// Must be UDP, TCP, or SCTP. Defaults to "TCP".
	Protocol Protocol `field:"optional" json:"protocol" yaml:"protocol"`
}

