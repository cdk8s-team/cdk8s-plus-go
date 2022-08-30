// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Emphemeral storage request and limit.
type EphemeralStorageResources struct {
	Limit cdk8s.Size `field:"optional" json:"limit" yaml:"limit"`
	Request cdk8s.Size `field:"optional" json:"request" yaml:"request"`
}

