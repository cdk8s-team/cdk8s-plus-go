// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22


// Options for `LabelSelector.of`.
type LabelSelectorOptions struct {
	// Expression based label matchers.
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Strict label matchers.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

