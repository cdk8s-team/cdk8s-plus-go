// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Memory request and limit.
type MemoryResources struct {
	Limit cdk8s.Size `field:"optional" json:"limit" yaml:"limit"`
	Request cdk8s.Size `field:"optional" json:"request" yaml:"request"`
}

