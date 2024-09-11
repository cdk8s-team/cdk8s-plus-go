package cdk8splus31

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `Namespace`.
type NamespaceProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
}

