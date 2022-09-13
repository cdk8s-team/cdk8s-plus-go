package k8s


// EndpointHints provides hints describing how an endpoint should be consumed.
type EndpointHintsV1Beta1 struct {
	// forZones indicates the zone(s) this endpoint should be consumed by to enable topology aware routing.
	//
	// May contain a maximum of 8 entries.
	ForZones *[]*ForZoneV1Beta1 `field:"optional" json:"forZones" yaml:"forZones"`
}

