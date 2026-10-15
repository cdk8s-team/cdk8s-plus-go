package k8s


// DeviceTaintRuleSpec specifies the selector and one taint.
type DeviceTaintRuleSpecV1Alpha3 struct {
	// The taint that gets applied to matching devices.
	Taint *DeviceTaintV1Alpha3 `field:"required" json:"taint" yaml:"taint"`
	// DeviceSelector defines which device(s) the taint is applied to.
	//
	// All selector criteria must be satisfied for a device to match. The empty selector matches all devices. Without a selector, no devices are matches.
	DeviceSelector *DeviceTaintSelectorV1Alpha3 `field:"optional" json:"deviceSelector" yaml:"deviceSelector"`
}

