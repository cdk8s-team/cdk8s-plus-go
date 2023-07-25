package cdk8splus27


// Options of `Volume.fromAwsElasticBlockStore`.
type AwsElasticBlockStoreVolumeOptions struct {
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// Specify "true" to force and set the ReadOnly property in VolumeMounts to "true".
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

