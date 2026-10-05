package k8s


// DeviceSelector must have exactly one field set.
type DeviceSelectorV1Beta2 struct {
	// CEL contains a CEL expression for selecting a device.
	Cel *CelDeviceSelectorV1Beta2 `field:"optional" json:"cel" yaml:"cel"`
}

