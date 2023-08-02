package cdk8splus25

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `AzureDiskPersistentVolume`.
type AzureDiskPersistentVolumeProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Contains all ways the volume can be mounted.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	//
	// Default: - No access modes.
	//
	AccessModes *[]PersistentVolumeAccessMode `field:"optional" json:"accessModes" yaml:"accessModes"`
	// Part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim.
	//
	// Expected to be non-nil when bound.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	//
	// Default: - Not bound to a specific claim.
	//
	Claim IPersistentVolumeClaim `field:"optional" json:"claim" yaml:"claim"`
	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	//
	// Default: - No options.
	//
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// When a user is done with their volume, they can delete the PVC objects from the API that allows reclamation of the resource.
	//
	// The reclaim policy tells the cluster what to do with
	// the volume after it has been released of its claim.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	//
	// Default: PersistentVolumeReclaimPolicy.RETAIN
	//
	ReclaimPolicy PersistentVolumeReclaimPolicy `field:"optional" json:"reclaimPolicy" yaml:"reclaimPolicy"`
	// What is the storage capacity of this volume.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	//
	// Default: - No specified.
	//
	Storage cdk8s.Size `field:"optional" json:"storage" yaml:"storage"`
	// Name of StorageClass to which this persistent volume belongs.
	// Default: - Volume does not belong to any storage class.
	//
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// Defines what type of volume is required by the claim.
	// Default: VolumeMode.FILE_SYSTEM
	//
	VolumeMode PersistentVolumeMode `field:"optional" json:"volumeMode" yaml:"volumeMode"`
	// The Name of the data disk in the blob storage.
	DiskName *string `field:"required" json:"diskName" yaml:"diskName"`
	// The URI the data disk in the blob storage.
	DiskUri *string `field:"required" json:"diskUri" yaml:"diskUri"`
	// Host Caching mode.
	// Default: - AzureDiskPersistentVolumeCachingMode.NONE.
	//
	CachingMode AzureDiskPersistentVolumeCachingMode `field:"optional" json:"cachingMode" yaml:"cachingMode"`
	// Filesystem type to mount.
	//
	// Must be a filesystem type supported by the host operating system.
	// Default: 'ext4'.
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// Kind of disk.
	// Default: AzureDiskPersistentVolumeKind.SHARED
	//
	Kind AzureDiskPersistentVolumeKind `field:"optional" json:"kind" yaml:"kind"`
	// Force the ReadOnly setting in VolumeMounts.
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

