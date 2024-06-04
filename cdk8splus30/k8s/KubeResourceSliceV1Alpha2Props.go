package k8s


// ResourceSlice provides information about available resources on individual nodes.
type KubeResourceSliceV1Alpha2Props struct {
	// DriverName identifies the DRA driver providing the capacity information.
	//
	// A field selector can be used to list only ResourceSlice objects with a certain driver name.
	DriverName *string `field:"required" json:"driverName" yaml:"driverName"`
	// Standard object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
	// NamedResources describes available resources using the named resources model.
	NamedResources *NamedResourcesResourcesV1Alpha2 `field:"optional" json:"namedResources" yaml:"namedResources"`
	// NodeName identifies the node which provides the resources if they are local to a node.
	//
	// A field selector can be used to list only ResourceSlice objects with a certain node name.
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
}

