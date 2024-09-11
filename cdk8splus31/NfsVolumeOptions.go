package cdk8splus31


// Options for the NFS based volume.
type NfsVolumeOptions struct {
	// Path that is exported by the NFS server.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Server is the hostname or IP address of the NFS server.
	Server *string `field:"required" json:"server" yaml:"server"`
	// If set to true, will force the NFS export to be mounted with read-only permissions.
	// Default: - false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

