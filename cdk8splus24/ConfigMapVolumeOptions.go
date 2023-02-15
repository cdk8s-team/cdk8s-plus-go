// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Options for the ConfigMap-based volume.
type ConfigMapVolumeOptions struct {
	// Mode bits to use on created files by default.
	//
	// Must be a value between 0 and
	// 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// If unspecified, each key-value pair in the Data field of the referenced ConfigMap will be projected into the volume as a file whose name is the key and content is the value.
	//
	// If specified, the listed keys will be projected
	// into the specified paths, and unlisted keys will not be present. If a key
	// is specified which is not present in the ConfigMap, the volume setup will
	// error unless it is marked optional. Paths must be relative and may not
	// contain the '..' path or start with '..'.
	Items *map[string]*PathMapping `field:"optional" json:"items" yaml:"items"`
	// The volume name.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the ConfigMap or its keys must be defined.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

