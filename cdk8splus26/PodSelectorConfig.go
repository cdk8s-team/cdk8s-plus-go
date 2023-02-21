// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


// Configuration for selecting pods, optionally in particular namespaces.
type PodSelectorConfig struct {
	// A selector to select pods by labels.
	LabelSelector LabelSelector `field:"required" json:"labelSelector" yaml:"labelSelector"`
	// Configuration for selecting which namepsaces are the pods allowed to be in.
	Namespaces *NamespaceSelectorConfig `field:"optional" json:"namespaces" yaml:"namespaces"`
}

