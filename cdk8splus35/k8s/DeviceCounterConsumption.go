package k8s


// DeviceCounterConsumption defines a set of counters that a device will consume from a CounterSet.
type DeviceCounterConsumption struct {
	// Counters defines the counters that will be consumed by the device.
	//
	// The maximum number of counters is 32.
	Counters *map[string]*Counter `field:"required" json:"counters" yaml:"counters"`
	// CounterSet is the name of the set from which the counters defined will be consumed.
	CounterSet *string `field:"required" json:"counterSet" yaml:"counterSet"`
}

