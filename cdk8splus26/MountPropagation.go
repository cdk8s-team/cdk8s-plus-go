// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


type MountPropagation string

const (
	// This volume mount will not receive any subsequent mounts that are mounted to this volume or any of its subdirectories by the host.
	//
	// In similar
	// fashion, no mounts created by the Container will be visible on the host.
	//
	// This is the default mode.
	//
	// This mode is equal to `private` mount propagation as described in the Linux
	// kernel documentation.
	MountPropagation_NONE MountPropagation = "NONE"
	// This volume mount will receive all subsequent mounts that are mounted to this volume or any of its subdirectories.
	//
	// In other words, if the host mounts anything inside the volume mount, the
	// Container will see it mounted there.
	//
	// Similarly, if any Pod with Bidirectional mount propagation to the same
	// volume mounts anything there, the Container with HostToContainer mount
	// propagation will see it.
	//
	// This mode is equal to `rslave` mount propagation as described in the Linux
	// kernel documentation.
	MountPropagation_HOST_TO_CONTAINER MountPropagation = "HOST_TO_CONTAINER"
	// This volume mount behaves the same the HostToContainer mount.
	//
	// In addition,
	// all volume mounts created by the Container will be propagated back to the
	// host and to all Containers of all Pods that use the same volume
	//
	// A typical use case for this mode is a Pod with a FlexVolume or CSI driver
	// or a Pod that needs to mount something on the host using a hostPath volume.
	//
	// This mode is equal to `rshared` mount propagation as described in the Linux
	// kernel documentation
	//
	// Caution: Bidirectional mount propagation can be dangerous. It can damage
	// the host operating system and therefore it is allowed only in privileged
	// Containers. Familiarity with Linux kernel behavior is strongly recommended.
	// In addition, any volume mounts created by Containers in Pods must be
	// destroyed (unmounted) by the Containers on termination.
	MountPropagation_BIDIRECTIONAL MountPropagation = "BIDIRECTIONAL"
)

