// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

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
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// The password or token for authentication.
	Password *string `field:"required" json:"password" yaml:"password"`
	// The user name for authentication.
	Username *string `field:"required" json:"username" yaml:"username"`
}

