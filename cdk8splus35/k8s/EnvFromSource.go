package k8s


// EnvFromSource represents the source of a set of ConfigMaps or Secrets.
type EnvFromSource struct {
	// The ConfigMap to select from.
	ConfigMapRef *ConfigMapEnvSource `field:"optional" json:"configMapRef" yaml:"configMapRef"`
	// Optional text to prepend to the name of each environment variable.
	//
	// May consist of any printable ASCII characters except '='.
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
	// The Secret to select from.
	SecretRef *SecretEnvSource `field:"optional" json:"secretRef" yaml:"secretRef"`
}

