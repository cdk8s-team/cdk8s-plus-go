package cdk8splus26


// Options for setting up backends for ingress rules.
type ServiceIngressBackendOptions struct {
	// The port to use to access the service.
	//
	// - This option will fail if the service does not expose any ports.
	// - If the service exposes multiple ports, this option must be specified.
	// - If the service exposes a single port, this option is optional and if
	//    specified, it must be the same port exposed by the service.
	// Default: - if the service exposes a single port, this port will be used.
	//
	Port *float64 `field:"optional" json:"port" yaml:"port"`
}

