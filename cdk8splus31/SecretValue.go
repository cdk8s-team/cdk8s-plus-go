package cdk8splus31


// Represents a specific value in JSON secret.
type SecretValue struct {
	// The JSON key.
	Key *string `field:"required" json:"key" yaml:"key"`
	// The secret.
	Secret ISecret `field:"required" json:"secret" yaml:"secret"`
}

