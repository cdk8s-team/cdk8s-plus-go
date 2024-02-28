package cdk8splus28

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `RoleBinding`.
type RoleBindingProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// The role to bind to.
	//
	// A RoleBinding can reference a Role or a ClusterRole.
	Role IRole `field:"required" json:"role" yaml:"role"`
}

