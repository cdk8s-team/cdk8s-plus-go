package k8s


// ParentReference describes a reference to a parent object.
type ParentReference struct {
	// name is the name of the object being referenced.
	Name *string `field:"required" json:"name" yaml:"name"`
	// resource is the resource of the object being referenced.
	Resource *string `field:"required" json:"resource" yaml:"resource"`
	// group is the group of the object being referenced.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// namespace is the namespace of the object being referenced.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
}

