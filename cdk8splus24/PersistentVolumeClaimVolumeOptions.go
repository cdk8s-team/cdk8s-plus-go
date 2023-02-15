// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Options for a PersistentVolumeClaim-based volume.
type PersistentVolumeClaimVolumeOptions struct {
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Will force the ReadOnly setting in VolumeMounts.
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

