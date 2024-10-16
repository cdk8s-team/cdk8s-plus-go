package cdk8splus30

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Options for `TlsSecret`.
type TlsSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	// Default: false.
	//
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// The TLS cert.
	TlsCert *string `field:"required" json:"tlsCert" yaml:"tlsCert"`
	// The TLS key.
	TlsKey *string `field:"required" json:"tlsKey" yaml:"tlsKey"`
}

