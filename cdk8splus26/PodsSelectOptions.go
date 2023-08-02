package cdk8splus26


// Options for `Pods.select`.
type PodsSelectOptions struct {
	// Expressions the pods must satisify.
	// Default: - no expressions requirements.
	//
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Labels the pods must have.
	// Default: - no strict labels requirements.
	//
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespaces the pods are allowed to be in.
	//
	// Use `Namespaces.all()` to allow all namespaces.
	// Default: - unset, implies the namespace of the resource this selection is used in.
	//
	Namespaces Namespaces `field:"optional" json:"namespaces" yaml:"namespaces"`
}

