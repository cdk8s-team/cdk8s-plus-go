package cdk8splus28


// Host path types.
type HostPathVolumeType string

const (
	// Empty string (default) is for backward compatibility, which means that no checks will be performed before mounting the hostPath volume.
	HostPathVolumeType_DEFAULT HostPathVolumeType = "DEFAULT"
	// If nothing exists at the given path, an empty directory will be created there as needed with permission set to 0755, having the same group and ownership with Kubelet.
	HostPathVolumeType_DIRECTORY_OR_CREATE HostPathVolumeType = "DIRECTORY_OR_CREATE"
	// A directory must exist at the given path.
	HostPathVolumeType_DIRECTORY HostPathVolumeType = "DIRECTORY"
	// If nothing exists at the given path, an empty file will be created there as needed with permission set to 0644, having the same group and ownership with Kubelet.
	HostPathVolumeType_FILE_OR_CREATE HostPathVolumeType = "FILE_OR_CREATE"
	// A file must exist at the given path.
	HostPathVolumeType_FILE HostPathVolumeType = "FILE"
	// A UNIX socket must exist at the given path.
	HostPathVolumeType_SOCKET HostPathVolumeType = "SOCKET"
	// A character device must exist at the given path.
	HostPathVolumeType_CHAR_DEVICE HostPathVolumeType = "CHAR_DEVICE"
	// A block device must exist at the given path.
	HostPathVolumeType_BLOCK_DEVICE HostPathVolumeType = "BLOCK_DEVICE"
)

