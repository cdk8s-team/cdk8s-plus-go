package cdk8splus31

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Options for `BasicAuthSecret`.
type BasicAuthSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	// Default: false.
	//
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// The password or token for authentication.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name for authentication.
	Username *string `field:"required" json:"username" yaml:"username"`
}

