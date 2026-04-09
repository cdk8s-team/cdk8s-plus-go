package k8s


// DeviceClassList is a collection of classes.
type KubeDeviceClassListProps struct {
	// Items is the list of resource classes.
	Items *[]*KubeDeviceClassProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

