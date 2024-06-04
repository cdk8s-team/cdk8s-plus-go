package k8s


// NamedResourcesInstance represents one individual hardware instance that can be selected based on its attributes.
type NamedResourcesInstanceV1Alpha2 struct {
	// Name is unique identifier among all resource instances managed by the driver on the node.
	//
	// It must be a DNS subdomain.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Attributes defines the attributes of this resource instance.
	//
	// The name of each attribute must be unique.
	Attributes *[]*NamedResourcesAttributeV1Alpha2 `field:"optional" json:"attributes" yaml:"attributes"`
}

