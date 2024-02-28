package cdk8splus28


// Options for `Pods.all`.
type PodsAllOptions struct {
	// Namespaces the pods are allowed to be in.
	//
	// Use `Namespaces.all()` to allow all namespaces.
	// Default: - unset, implies the namespace of the resource this selection is used in.
	//
	Namespaces Namespaces `field:"optional" json:"namespaces" yaml:"namespaces"`
}

