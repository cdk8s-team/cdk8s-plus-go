// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Azure Disk kinds.
type AzureDiskPersistentVolumeKind string

const (
	// Multiple blob disks per storage account.
	AzureDiskPersistentVolumeKind_SHARED AzureDiskPersistentVolumeKind = "SHARED"
	// Single blob disk per storage account.
	AzureDiskPersistentVolumeKind_DEDICATED AzureDiskPersistentVolumeKind = "DEDICATED"
	// Azure managed data disk.
	AzureDiskPersistentVolumeKind_MANAGED AzureDiskPersistentVolumeKind = "MANAGED"
)

