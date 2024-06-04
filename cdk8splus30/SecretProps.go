package cdk8splus30

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Options for `Secret`.
type SecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	// Default: false.
	//
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// stringData allows specifying non-binary secret data in string form.
	//
	// It is
	// provided as a write-only convenience method. All keys and values are merged
	// into the data field on write, overwriting any existing values. It is never
	// output when reading from the API.
	StringData *map[string]*string `field:"optional" json:"stringData" yaml:"stringData"`
	// Optional type associated with the secret.
	//
	// Used to facilitate programmatic
	// handling of secret data by various controllers.
	// Default: undefined - Don't set a type.
	//
	Type *string `field:"optional" json:"type" yaml:"type"`
}

