package k8s


// TypedLocalObjectReference allows to reference typed object inside the same namespace.
type TypedLocalObjectReferenceV1Alpha1 struct {
	// Kind is the type of resource being referenced.
	//
	// It must be a path segment name.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Name is the name of resource being referenced.
	//
	// It must be a path segment name.
	Name *string `field:"required" json:"name" yaml:"name"`
	// APIGroup is the group for the resource being referenced.
	//
	// If APIGroup is empty, the specified Kind must be in the core API group. For any other third-party types, setting APIGroup is required. It must be a DNS subdomain.
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
}

