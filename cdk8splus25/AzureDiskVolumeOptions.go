package cdk8splus25


// Options of `Volume.fromAzureDisk`.
type AzureDiskVolumeOptions struct {
	// Host Caching mode.
	// Default: - AzureDiskPersistentVolumeCachingMode.NONE.
	//
	CachingMode AzureDiskPersistentVolumeCachingMode `field:"optional" json:"cachingMode" yaml:"cachingMode"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	// Default: 'ext4'.
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Kind of disk.
	// Default: AzureDiskPersistentVolumeKind.SHARED
	//
	Kind AzureDiskPersistentVolumeKind `field:"optional" json:"kind" yaml:"kind"`
	// The volume name.
	// Default: - auto-generated.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Force the ReadOnly setting in VolumeMounts.
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

