package cdk8splus27


// Properties for `PodSecurityContext`.
type PodSecurityContextProps struct {
	// Indicates that the container must run as a non-root user.
	//
	// If true, the Kubelet will validate the image at runtime to ensure that it does
	// not run as UID 0 (root) and fail to start the container if it does.
	// Default: true.
	//
	EnsureNonRoot *bool `field:"optional" json:"ensureNonRoot" yaml:"ensureNonRoot"`
	// Modify the ownership and permissions of pod volumes to this GID.
	// Default: - Volume ownership is not changed.
	//
	FsGroup *float64 `field:"optional" json:"fsGroup" yaml:"fsGroup"`
	// Defines behavior of changing ownership and permission of the volume before being exposed inside Pod.
	//
	// This field will only apply to volume types which support fsGroup based ownership(and permissions).
	// It will have no effect on ephemeral volume types such as: secret, configmaps and emptydir.
	// Default: FsGroupChangePolicy.ALWAYS
	//
	FsGroupChangePolicy FsGroupChangePolicy `field:"optional" json:"fsGroupChangePolicy" yaml:"fsGroupChangePolicy"`
	// The GID to run the entrypoint of the container process.
	// Default: - Group configured by container runtime.
	//
	Group *float64 `field:"optional" json:"group" yaml:"group"`
	// Sysctls hold a list of namespaced sysctls used for the pod.
	//
	// Pods with unsupported sysctls (by the container runtime) might fail to launch.
	// Default: - No sysctls.
	//
	Sysctls *[]*Sysctl `field:"optional" json:"sysctls" yaml:"sysctls"`
	// The UID to run the entrypoint of the container process.
	// Default: - User specified in image metadata.
	//
	User *float64 `field:"optional" json:"user" yaml:"user"`
}

