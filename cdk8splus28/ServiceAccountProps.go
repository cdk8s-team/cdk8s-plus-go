package cdk8splus28

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Properties for initialization of `ServiceAccount`.
type ServiceAccountProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// Indicates whether pods running as this service account should have an API token automatically mounted.
	//
	// Can be overridden at the pod level.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/#use-the-default-service-account-to-access-the-api-server
	//
	// Default: false.
	//
	AutomountToken *bool `field:"optional" json:"automountToken" yaml:"automountToken"`
	// List of secrets allowed to be used by pods running using this ServiceAccount.
	// See: https://kubernetes.io/docs/concepts/configuration/secret
	//
	Secrets *[]ISecret `field:"optional" json:"secrets" yaml:"secrets"`
}

