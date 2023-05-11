package cdk8splus26


// Options for `Pods.all`.
type PodsAllOptions struct {
	// Namespaces the pods are allowed to be in.
	//
	// Use `Namespaces.all()` to allow all namespaces.
	Namespaces Namespaces `field:"optional" json:"namespaces" yaml:"namespaces"`
}

