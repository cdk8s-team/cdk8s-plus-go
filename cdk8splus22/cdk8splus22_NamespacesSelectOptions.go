// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Options for `Namespaces.select`.
type NamespacesSelectOptions struct {
	// Namespaces must satisfy these selectors.
	//
	// The selectors query labels, just like the `labels` property, but they
	// provide a more advanced matching mechanism.
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Labels the namespaces must have.
	//
	// This is equivalent to using an 'Is' selector.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespaces names must be one of these.
	Names *[]*string `field:"optional" json:"names" yaml:"names"`
}

