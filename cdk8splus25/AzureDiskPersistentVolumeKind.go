// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


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

