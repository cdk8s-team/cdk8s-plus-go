// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Options for the CSI driver based volume.
type CsiVolumeOptions struct {
	// Any driver-specific attributes to pass to the CSI volume builder.
	Attributes *map[string]*string `field:"optional" json:"attributes" yaml:"attributes"`
	// The filesystem type to mount.
	//
	// Ex. "ext4", "xfs", "ntfs". If not provided,
	// the empty value is passed to the associated CSI driver, which will
	// determine the default filesystem to apply.
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Whether the mounted volume should be read-only or not.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

