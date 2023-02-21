// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26


// Options for `Pods.select`.
type PodsSelectOptions struct {
	// Expressions the pods must satisify.
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Labels the pods must have.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespaces the pods are allowed to be in.
	//
	// Use `Namespaces.all()` to allow all namespaces.
	Namespaces Namespaces `field:"optional" json:"namespaces" yaml:"namespaces"`
}

