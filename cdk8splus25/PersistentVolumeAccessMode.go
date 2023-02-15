// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Access Modes.
type PersistentVolumeAccessMode string

const (
	// The volume can be mounted as read-write by a single node.
	//
	// ReadWriteOnce access mode still can allow multiple pods to access
	// the volume when the pods are running on the same node.
	PersistentVolumeAccessMode_READ_WRITE_ONCE PersistentVolumeAccessMode = "READ_WRITE_ONCE"
	// The volume can be mounted as read-only by many nodes.
	PersistentVolumeAccessMode_READ_ONLY_MANY PersistentVolumeAccessMode = "READ_ONLY_MANY"
	// The volume can be mounted as read-write by many nodes.
	PersistentVolumeAccessMode_READ_WRITE_MANY PersistentVolumeAccessMode = "READ_WRITE_MANY"
	// The volume can be mounted as read-write by a single Pod.
	//
	// Use ReadWriteOncePod access mode if you want to ensure that
	// only one pod across whole cluster can read that PVC or write to it.
	// This is only supported for CSI volumes and Kubernetes version 1.22+.
	PersistentVolumeAccessMode_READ_WRITE_ONCE_POD PersistentVolumeAccessMode = "READ_WRITE_ONCE_POD"
)

