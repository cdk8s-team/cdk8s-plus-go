package cdk8splus26


// Options for the Secret-based volume.
type SecretVolumeOptions struct {
	// Mode bits to use on created files by default.
	//
	// Must be a value between 0 and
	// 0777. Defaults to 0644. Directories within the path are not affected by
	// this setting. This might be in conflict with other options that affect the
	// file mode, like fsGroup, and the result can be other mode bits set.
	// Default: 0644. Directories within the path are not affected by this
	// setting. This might be in conflict with other options that affect the file
	// mode, like fsGroup, and the result can be other mode bits set.
	//
	DefaultMode *float64 `field:"optional" json:"defaultMode" yaml:"defaultMode"`
	// If unspecified, each key-value pair in the Data field of the referenced secret will be projected into the volume as a file whose name is the key and content is the value.
	//
	// If specified, the listed keys will be projected
	// into the specified paths, and unlisted keys will not be present. If a key
	// is specified which is not present in the secret, the volume setup will
	// error unless it is marked optional. Paths must be relative and may not
	// contain the '..' path or start with '..'.
	// Default: - no mapping.
	//
	Items *map[string]*PathMapping `field:"optional" json:"items" yaml:"items"`
	// The volume name.
	// Default: - auto-generated.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specify whether the secret or its keys must be defined.
	// Default: - undocumented.
	//
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

