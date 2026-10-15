package k8s


// BoundObjectReference is a reference to an object that a token is bound to.
type BoundObjectReference struct {
	// apiVersion is API version of the referent.
	ApiVersion *string `field:"optional" json:"apiVersion" yaml:"apiVersion"`
	// kind of the referent.
	//
	// Valid kinds are 'Pod' and 'Secret'.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// name of the referent.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// uid of the referent.
	Uid *string `field:"optional" json:"uid" yaml:"uid"`
}

