package k8s


// Counter describes a quantity associated with a device.
type CounterV1Beta2 struct {
	// Value defines how much of a certain device counter is available.
	Value Quantity `field:"required" json:"value" yaml:"value"`
}

