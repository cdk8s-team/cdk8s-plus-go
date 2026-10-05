package k8s


// DeviceClassConfiguration is used in DeviceClass.
type DeviceClassConfigurationV1Beta2 struct {
	// Opaque provides driver-specific configuration parameters.
	Opaque *OpaqueDeviceConfigurationV1Beta2 `field:"optional" json:"opaque" yaml:"opaque"`
}

