package k8s


// GRPCAction specifies an action involving a GRPC service.
type GrpcAction struct {
	// Port number of the gRPC service.
	//
	// Number must be in the range 1 to 65535.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// mode specifies the connection mode for the gRPC health probe.
	//
	// Set to "TLS" to use TLS without certificate verification. Set to "Plaintext" to use a plaintext (insecure) connection explicitly. If not specified, the probe uses a plaintext (insecure) connection.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Service is the name of the service to place in the gRPC HealthCheckRequest (see https://github.com/grpc/grpc/blob/master/doc/health-checking.md).
	//
	// If this is not specified, the default behavior is defined by gRPC.
	Service *string `field:"optional" json:"service" yaml:"service"`
}

