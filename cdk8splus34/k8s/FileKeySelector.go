package k8s


// FileKeySelector selects a key of the env file.
type FileKeySelector struct {
	// The key within the env file.
	//
	// An invalid key will prevent the pod from starting. The keys defined within a source may consist of any printable ASCII characters except '='. During Alpha stage of the EnvFiles feature gate, the key size is limited to 128 characters.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The path within the volume from which to select the file.
	//
	// Must be relative and may not contain the '..' path or start with '..'.
	Path *string `field:"required" json:"path" yaml:"path"`
	// The name of the volume mount containing the env file.
	VolumeName *string `field:"required" json:"volumeName" yaml:"volumeName"`
	// Specify whether the file or its key must be defined.
	//
	// If the file or key does not exist, then the env var is not published. If optional is set to true and the specified key does not exist, the environment variable will not be set in the Pod's containers.
	//
	// If optional is set to false and the specified key does not exist, an error will be returned during Pod creation.
	Optional *bool `field:"optional" json:"optional" yaml:"optional"`
}

