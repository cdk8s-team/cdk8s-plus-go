package k8s


// DeviceClassConfiguration is used in DeviceClass.
type DeviceClassConfiguration struct {
	// Opaque provides driver-specific configuration parameters.
	Opaque *OpaqueDeviceConfiguration `field:"optional" json:"opaque" yaml:"opaque"`
}

