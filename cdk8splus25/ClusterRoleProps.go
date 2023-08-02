package cdk8splus25

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `ClusterRole`.
type ClusterRoleProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Specify labels that should be used to locate ClusterRoles, whose rules will be automatically filled into this ClusterRole's rules.
	AggregationLabels *map[string]*string `field:"optional" json:"aggregationLabels" yaml:"aggregationLabels"`
	// A list of rules the role should allow.
	// Default: [].
	//
	Rules *[]*ClusterRolePolicyRule `field:"optional" json:"rules" yaml:"rules"`
}

