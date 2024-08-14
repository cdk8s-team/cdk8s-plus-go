package cdk8splus29


// Properties for `ContainerSecurityContext`.
type ContainerSecurityContextProps struct {
	// Whether a process can gain more privileges than its parent process.
	// Default: false.
	//
	AllowPrivilegeEscalation *bool `field:"optional" json:"allowPrivilegeEscalation" yaml:"allowPrivilegeEscalation"`
	// POSIX capabilities for running containers.
	// Default: none.
	//
	Capabilities *ContainerSecutiryContextCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it does
	// not run as UID 0 (root) and fail to start the container if it does.
	// Default: true.
	//
	EnsureNonRoot *bool `field:"optional" json:"ensureNonRoot" yaml:"ensureNonRoot"`
	// The GID to run the entrypoint of the container process.
	// Default: - 26000. An arbitrary number bigger than 9999 is selected here.
	// This is so that the container is blocked to access host files even if
	// somehow it manages to get access to host file system.
	//
	Group *float64 `field:"optional" json:"group" yaml:"group"`
	// Run container in privileged mode.
	//
	// Processes in privileged containers are essentially equivalent to root on the host.
	// Default: false.
	//
	Privileged *bool `field:"optional" json:"privileged" yaml:"privileged"`
	// Whether this container has a read-only root filesystem.
	// Default: true.
	//
	ReadOnlyRootFilesystem *bool `field:"optional" json:"readOnlyRootFilesystem" yaml:"readOnlyRootFilesystem"`
	// Container's seccomp profile settings.
	//
	// Only one profile source may be set.
	// Default: none.
	//
	SeccompProfile *SeccompProfile `field:"optional" json:"seccompProfile" yaml:"seccompProfile"`
	// The UID to run the entrypoint of the container process.
	// Default: - 25000. An arbitrary number bigger than 9999 is selected here.
	// This is so that the container is blocked to access host files even if
	// somehow it manages to get access to host file system.
	//
	User *float64 `field:"optional" json:"user" yaml:"user"`
}

