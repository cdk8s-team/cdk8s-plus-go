package cdk8splus27

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `GCEPersistentDiskPersistentVolume`.
type GCEPersistentDiskPersistentVolumeProps struct {
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
	// Unique name of the PD resource in GCE.
	//
	// Used to identify the disk in GCE.
	// See: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	//
	PdName *string `field:"required" json:"pdName" yaml:"pdName"`
	// Filesystem type of the volume that you want to mount.
	//
	// Tip: Ensure that the filesystem type is supported by the host operating system.
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	// Default: 'ext4'.
	//
	FsType *string `field:"optional" json:"fsType" yaml:"fsType"`
	// The partition in the volume that you want to mount.
	//
	// If omitted, the default is to mount by volume name.
	// Examples: For volume /dev/sda1, you specify the partition as "1".
	// Similarly, the volume partition for /dev/sda is "0" (or you can leave the property empty).
	// Default: - No partition.
	//
	Partition *float64 `field:"optional" json:"partition" yaml:"partition"`
	// Specify "true" to force and set the ReadOnly property in VolumeMounts to "true".
	// See: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	//
	// Default: false.
	//
	ReadOnly *bool `field:"optional" json:"readOnly" yaml:"readOnly"`
}

