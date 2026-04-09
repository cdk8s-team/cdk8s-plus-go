package k8s


// EndpointHints provides hints describing how an endpoint should be consumed.
type EndpointHints struct {
	// forNodes indicates the node(s) this endpoint should be consumed by when using topology aware routing.
	//
	// May contain a maximum of 8 entries. This is an Alpha feature and is only used when the PreferSameTrafficDistribution feature gate is enabled.
	ForNodes *[]*ForNode `field:"optional" json:"forNodes" yaml:"forNodes"`
	// forZones indicates the zone(s) this endpoint should be consumed by when using topology aware routing.
	//
	// May contain a maximum of 8 entries.
	ForZones *[]*ForZone `field:"optional" json:"forZones" yaml:"forZones"`
}

