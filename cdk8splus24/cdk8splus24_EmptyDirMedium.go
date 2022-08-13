// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// The medium on which to store the volume.
type EmptyDirMedium string

const (
	// The default volume of the backing node.
	EmptyDirMedium_DEFAULT EmptyDirMedium = "DEFAULT"
	// Mount a tmpfs (RAM-backed filesystem) for you instead.
	//
	// While tmpfs is very
	// fast, be aware that unlike disks, tmpfs is cleared on node reboot and any
	// files you write will count against your Container's memory limit.
	EmptyDirMedium_MEMORY EmptyDirMedium = "MEMORY"
)

