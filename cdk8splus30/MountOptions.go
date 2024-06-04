package cdk8splus30


// Options for mounts.
type MountOptions struct {
	// Determines how mounts are propagated from the host to container and the other way around.
	//
	// When not set, MountPropagationNone is used.
	//
	// Mount propagation allows for sharing volumes mounted by a Container to
	// other Containers in the same Pod, or even to other Pods on the same node.
	// Default: MountPropagation.NONE
	//
	Propagation MountPropagation `field:"optional" json:"propagation" yaml:"propagation"`
	// Mounted read-only if true, read-write otherwise (false or unspecified).
	//
	// Defaults to false.
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Path within the volume from which the container's volume should be mounted.).
	// Default: "" the volume's root.
	//
	SubPath *string `field:"optional" json:"subPath" yaml:"subPath"`
	// Expanded path within the volume from which the container's volume should be mounted.
	//
	// Behaves similarly to SubPath but environment variable references
	// $(VAR_NAME) are expanded using the container's environment. Defaults to ""
	// (volume's root).
	//
	// `subPathExpr` and `subPath` are mutually exclusive.
	// Default: "" volume's root.
	//
	SubPathExpr *string `field:"optional" json:"subPathExpr" yaml:"subPathExpr"`
}

