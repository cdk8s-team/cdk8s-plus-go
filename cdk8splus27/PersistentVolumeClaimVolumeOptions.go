package cdk8splus27


// Options for a PersistentVolumeClaim-based volume.
type PersistentVolumeClaimVolumeOptions struct {
	// The volume name.
	// Default: - Derived from the PVC name.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Will force the ReadOnly setting in VolumeMounts.
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

