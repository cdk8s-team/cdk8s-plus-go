package cdk8splus29


// Options for `Handler.fromTcpSocket`.
type HandlerFromTcpSocketOptions struct {
	// The host name to connect to on the container.
	// Default: - defaults to the pod IP.
	//
	Host *string `field:"optional" json:"host" yaml:"host"`
	// The TCP port to connect to on the container.
	// Default: - defaults to `container.port`.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

