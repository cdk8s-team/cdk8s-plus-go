package cdk8splus26


// Properties for `ContainerSecurityContext`.
type ContainerSecurityContextProps struct {
	// Whether a process can gain more privileges than its parent process.
	AllowPrivilegeEscalation *bool `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it does
	// not run as UID 0 (root) and fail to start the container if it does.
	EnsureNonRoot *bool `field:"optional" json:"ensureNonRoot" yaml:"ensureNonRoot"`
	// The GID to run the entrypoint of the container process.
	Group *float64 `field:"optional" json:"group" yaml:"group"`
	// Run container in privileged mode.
	//
	// Processes in privileged containers are essentially equivalent to root on the host.
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// Whether this container has a read-only root filesystem.
	ReadOnlyRootFilesystem *bool `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// The UID to run the entrypoint of the container process.
	User *float64 `field:"optional" json:"user" yaml:"user"`
}

