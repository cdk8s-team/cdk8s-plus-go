package cdk8splus25


// Options for the CSI driver based volume.
type CsiVolumeOptions struct {
	// Any driver-specific attributes to pass to the CSI volume builder.
	// Default: - undefined.
	//
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// The filesystem type to mount.
	//
	// Ex. "ext4", "xfs", "ntfs". If not provided,
	// the empty value is passed to the associated CSI driver, which will
	// determine the default filesystem to apply.
	// Default: - driver-dependent.
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The volume name.
	// Default: - auto-generated.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether the mounted volume should be read-only or not.
	// Default: - false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

