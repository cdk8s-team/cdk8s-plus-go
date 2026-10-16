package k8s


// RoleRef contains information that points to the role being used.
type RoleRef struct {
	// kind is the type of resource being referenced.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// name is the name of resource being referenced.
	Name *string `field:"required" json:"name" yaml:"name"`
	// apiGroup is the group for the resource being referenced.
	ApiGroup *string `field:"optional" json:"apiGroup" yaml:"apiGroup"`
}

