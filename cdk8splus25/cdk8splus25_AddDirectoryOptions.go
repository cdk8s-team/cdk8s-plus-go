// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25


// Options for `configmap.addDirectory()`.
type AddDirectoryOptions struct {
	// Glob patterns to exclude when adding files.
	Exclude *[]*string `field:"optional" json:"exclude" yaml:"exclude"`
	// A prefix to add to all keys in the config map.
	KeyPrefix *string `field:"optional" json:"keyPrefix" yaml:"keyPrefix"`
}

