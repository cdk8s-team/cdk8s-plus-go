package k8s


// ResourceClaimSpec defines what is being requested in a ResourceClaim and how to configure it.
type ResourceClaimSpec struct {
	// Devices defines how to request devices.
	Devices *DeviceClaim `field:"optional" json:"devices" yaml:"devices"`
}

