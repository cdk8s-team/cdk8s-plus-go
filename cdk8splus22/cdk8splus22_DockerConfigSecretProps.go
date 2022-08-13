// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Options for `DockerConfigSecret`.
type DockerConfigSecretProps struct {
	// Metadata that all persisted resources must have, which includes all objects users must create.
	Metadata *cdk8s.ApiObjectMetadata `field:"optional" json:"metadata" yaml:"metadata"`
	// If set to true, ensures that data stored in the Secret cannot be updated (only object metadata can be modified).
	//
	// If not set to true, the field can be modified at any time.
	Immutable *bool `field:"optional" json:"immutable" yaml:"immutable"`
	// JSON content to provide for the `~/.docker/config.json` file. This will be stringified and inserted as stringData.
	// See: https://docs.docker.com/engine/reference/commandline/cli/#sample-configuration-file
	//
	Data *map[string]interface{} `field:"required" json:"data" yaml:"data"`
}

