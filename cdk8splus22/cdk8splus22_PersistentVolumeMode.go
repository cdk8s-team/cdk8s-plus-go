// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Volume Modes.
type PersistentVolumeMode string

const (
	// Volume is ounted into Pods into a directory.
	//
	// If the volume is backed by a block device and the device is empty,
	// Kubernetes creates a filesystem on the device before mounting it
	// for the first time.
	PersistentVolumeMode_FILE_SYSTEM PersistentVolumeMode = "FILE_SYSTEM"
	// Use a volume as a raw block device.
	//
	// Such volume is presented into a Pod as a block device,
	// without any filesystem on it. This mode is useful to provide a Pod the fastest possible way
	// to access a volume, without any filesystem layer between the Pod
	// and the volume. On the other hand, the application running in
	// the Pod must know how to handle a raw block device.
	PersistentVolumeMode_BLOCK PersistentVolumeMode = "BLOCK"
)

