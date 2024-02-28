package cdk8splus28


// Configuration for selecting namespaces.
type NamespaceSelectorConfig struct {
	// A selector to select namespaces by labels.
	LabelSelector LabelSelector `field:"optional" json:"labelSelector" yaml:"labelSelector"`
	// A list of names to select namespaces by names.
	Names *[]*string `field:"optional" json:"names" yaml:"names"`
}

