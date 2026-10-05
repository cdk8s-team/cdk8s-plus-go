package k8s


// ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
type ResourceClaimSpecV1Beta2 struct {
	// Devices defines how to request devices.
	Devices *DeviceClaimV1Beta2 `field:"optional" json:"devices" yaml:"devices"`
}

