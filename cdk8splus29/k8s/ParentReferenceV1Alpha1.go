package k8s


// ParentReference describes a reference to a parent object.
type ParentReferenceV1Alpha1 struct {
	// Group is the group of the object being referenced.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Name is the name of the object being referenced.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Namespace is the namespace of the object being referenced.
	Namespace *string `field:"optional" json:"namespace" yaml:"namespace"`
	// Resource is the resource of the object being referenced.
	Resource *string `field:"optional" json:"resource" yaml:"resource"`
}

