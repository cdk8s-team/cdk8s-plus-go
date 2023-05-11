package cdk8splus26


// Options of `Volume.fromAzureDisk`.
type AzureDiskVolumeOptions struct {
	// Host Caching mode.
	CachingMode AzureDiskPersistentVolumeCachingMode `field:"optional" json:"cachingMode" yaml:"cachingMode"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Kind of disk.
	Kind AzureDiskPersistentVolumeKind `field:"optional" json:"kind" yaml:"kind"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

