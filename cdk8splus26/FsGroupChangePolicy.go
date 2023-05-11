package cdk8splus26


type FsGroupChangePolicy string

const (
	// Only change permissions and ownership if permission and ownership of root directory does not match with expected permissions of the volume.
	//
	// This could help shorten the time it takes to change ownership and permission of a volume.
	FsGroupChangePolicy_ON_ROOT_MISMATCH FsGroupChangePolicy = "ON_ROOT_MISMATCH"
	// Always change permission and ownership of the volume when volume is mounted.
	FsGroupChangePolicy_ALWAYS FsGroupChangePolicy = "ALWAYS"
)

