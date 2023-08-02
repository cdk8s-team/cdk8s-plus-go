package cdk8splus27


// Options for a HostPathVolume-based volume.
type HostPathVolumeOptions struct {
	// The path of the directory on the host.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The expected type of the path found on the host.
	// Default: HostPathVolumeType.DEFAULT
	//
	Type HostPathVolumeType `field:"optional" json:"type" yaml:"type"`
}

