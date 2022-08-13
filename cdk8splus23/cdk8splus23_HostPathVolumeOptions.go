// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23


// Options for a HostPathVolume-based volume.
type HostPathVolumeOptions struct {
	// The path of the directory on the host.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The expected type of the path found on the host.
	Type HostPathVolumeType `field:"optional" json:"type" yaml:"type"`
}

