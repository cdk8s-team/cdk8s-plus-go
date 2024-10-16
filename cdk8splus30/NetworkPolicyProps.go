package cdk8splus30

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for `NetworkPolicy`.
type NetworkPolicyProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Egress traffic configuration.
	// Default: - the policy doesn't change egress behavior of the pods it selects.
	//
	Egress *NetworkPolicyTraffic `field:"optional" json:"egress" yaml:"egress"`
	// Ingress traffic configuration.
	// Default: - the policy doesn't change ingress behavior of the pods it selects.
	//
	Ingress *NetworkPolicyTraffic `field:"optional" json:"ingress" yaml:"ingress"`
	// Which pods does this policy object applies to.
	//
	// This can either be a single pod / workload, or a grouping of pods selected
	// via the `Pods.select` function. Rules is applied to any pods selected by this property.
	// Multiple network policies can select the same set of pods.
	// In this case, the rules for each are combined additively.
	//
	// Note that.
	// Default: - will select all pods in the namespace of the policy.
	//
	Selector IPodSelector `field:"optional" json:"selector" yaml:"selector"`
}

