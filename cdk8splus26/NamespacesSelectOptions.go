package cdk8splus26


// Options for `Namespaces.select`.
type NamespacesSelectOptions struct {
	// Namespaces must satisfy these selectors.
	//
	// The selectors query labels, just like the `labels` property, but they
	// provide a more advanced matching mechanism.
	// Default: - no selector requirements.
	//
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Labels the namespaces must have.
	//
	// This is equivalent to using an 'Is' selector.
	// Default: - no strict labels requirements.
	//
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// Namespaces names must be one of these.
	// Default: - no name requirements.
	//
	Names *[]*string `field:"optional" json:"names" yaml:"names"`
}

