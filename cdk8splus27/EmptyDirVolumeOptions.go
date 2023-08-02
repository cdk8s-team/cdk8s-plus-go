package cdk8splus27

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Options for volumes populated with an empty directory.
type EmptyDirVolumeOptions struct {
	// By default, emptyDir volumes are stored on whatever medium is backing the node - that might be disk or SSD or network storage, depending on your environment.
	//
	// However, you can set the emptyDir.medium field to
	// `EmptyDirMedium.MEMORY` to tell Kubernetes to mount a tmpfs (RAM-backed
	// filesystem) for you instead. While tmpfs is very fast, be aware that unlike
	// disks, tmpfs is cleared on node reboot and any files you write will count
	// against your Container's memory limit.
	// Default: EmptyDirMedium.DEFAULT
	//
	Medium EmptyDirMedium `field:"optional" json:"medium" yaml:"medium"`
	// Total amount of local storage required for this EmptyDir volume.
	//
	// The size
	// limit is also applicable for memory medium. The maximum usage on memory
	// medium EmptyDir would be the minimum value between the SizeLimit specified
	// here and the sum of memory limits of all containers in a pod.
	// Default: - limit is undefined.
	//
	SizeLimit cdk8s.Size `field:"optional" json:"sizeLimit" yaml:"sizeLimit"`
}

