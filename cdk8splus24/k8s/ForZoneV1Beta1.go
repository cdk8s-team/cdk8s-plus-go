package k8s


// ForZone provides information about which zones should consume this endpoint.
type ForZoneV1Beta1 struct {
	// name represents the name of the zone.
	Name *string `field:"required" json:"name" yaml:"name"`
}

