// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Azure disk caching modes.
type AzureDiskPersistentVolumeCachingMode string

const (
	// None.
	AzureDiskPersistentVolumeCachingMode_NONE AzureDiskPersistentVolumeCachingMode = "NONE"
	// ReadOnly.
	AzureDiskPersistentVolumeCachingMode_READ_ONLY AzureDiskPersistentVolumeCachingMode = "READ_ONLY"
	// ReadWrite.
	AzureDiskPersistentVolumeCachingMode_READ_WRITE AzureDiskPersistentVolumeCachingMode = "READ_WRITE"
)

