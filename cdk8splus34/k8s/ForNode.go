package k8s


// ForNode provides information about which nodes should consume this endpoint.
type ForNode struct {
	// name represents the name of the node.
	Name *string `field:"required" json:"name" yaml:"name"`
}

