package cdk8splus29


// Options for `LabelSelector.of`.
type LabelSelectorOptions struct {
	// Expression based label matchers.
	Expressions *[]LabelExpression `field:"optional" json:"expressions" yaml:"expressions"`
	// Strict label matchers.
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
}

