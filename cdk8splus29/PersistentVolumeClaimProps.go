package cdk8splus29

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `PersistentVolumeClaim`.
type PersistentVolumeClaimProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Contains the access modes the volume should support.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes-1
	//
	// Default: - No access modes requirement.
	//
	AccessModes *[]PersistentVolumeAccessMode `field:"optional" json:"accessModes" yaml:"accessModes"`
	// Minimum storage size the volume should have.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#resources
	//
	// Default: - No storage requirement.
	//
	Storage cdk8s.Size `field:"optional" json:"storage" yaml:"storage"`
	// Name of the StorageClass required by the claim. When this property is not set, the behavior is as follows:.
	//
	// - If the admission plugin is turned on, the storage class marked as default will be used.
	// - If the admission plugin is turned off, the pvc can only be bound to volumes without a storage class.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes#class-1
	//
	// Default: - Not set.
	//
	StorageClassName *string `field:"optional" json:"storageClassName" yaml:"storageClassName"`
	// The PersistentVolume backing this claim.
	//
	// The control plane still checks that storage class, access modes,
	// and requested storage size on the volume are valid.
	//
	// Note that in order to guarantee a proper binding, the volume should
	// also define a `claimRef` referring to this claim. Otherwise, the volume may be
	// claimed be other pvc's before it gets a chance to bind to this one.
	//
	// If the volume is managed (i.e not imported), you can use `pv.claim()` to easily
	// create a bi-directional bounded claim.
	// See: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#binding.
	//
	// Default: - No specific volume binding.
	//
	Volume IPersistentVolume `field:"optional" json:"volume" yaml:"volume"`
	// Defines what type of volume is required by the claim.
	// Default: VolumeMode.FILE_SYSTEM
	//
	VolumeMode PersistentVolumeMode `field:"optional" json:"volumeMode" yaml:"volumeMode"`
}

