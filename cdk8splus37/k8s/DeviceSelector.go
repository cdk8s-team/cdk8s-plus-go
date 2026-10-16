package k8s


// DeviceSelector must have exactly one field set.
type DeviceSelector struct {
	// CEL contains a CEL expression for selecting a device.
	Cel *CelDeviceSelector `field:"optional" json:"cel" yaml:"cel"`
}

