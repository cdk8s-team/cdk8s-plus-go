package cdk8splus25

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Memory request and limit.
type MemoryResources struct {
	Limit cdk8s.Size `field:"optional" json:"limit" yaml:"limit"`
	Request cdk8s.Size `field:"optional" json:"request" yaml:"request"`
}

