// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Configuration for selecting namespaces.
type NamespaceSelectorConfig struct {
	// A selector to select namespaces by labels.
	LabelSelector LabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// A list of names to select namespaces by names.
	Names *[]*string `field:"optional" json:"names" yaml:"names"`
}

