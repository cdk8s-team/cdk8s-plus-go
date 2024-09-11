package cdk8splus31


type SeccompProfile struct {
	// Indicates which kind of seccomp profile will be applied.
	Type SeccompProfileType `field:"required" json:"type" yaml:"type"`
	// localhostProfile indicates a profile defined in a file on the node should be used.
	//
	// The profile must be preconfigured on the node to work. Must be a descending path,
	// relative to the kubelet's configured seccomp profile location.
	// Must only be set if type is "Localhost".
	// Default: - empty string.
	//
	LocalhostProfile *string `field:"optional" json:"localhostProfile" yaml:"localhostProfile"`
}

