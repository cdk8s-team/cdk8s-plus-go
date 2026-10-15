package k8s


// CounterSet defines a named set of counters that are available to be used by devices defined in the ResourcePool.
//
// The counters are not allocatable by themselves, but can be referenced by devices. When a device is allocated, the portion of counters it uses will no longer be available for use by other devices.
type CounterSet struct {
	// Counters defines the set of counters for this CounterSet The name of each counter must be unique in that set and must be a DNS label.
	//
	// The maximum number of counters is 32.
	Counters *map[string]*Counter `field:"required" json:"counters" yaml:"counters"`
	// Name defines the name of the counter set.
	//
	// It must be a DNS label.
	Name *string `field:"required" json:"name" yaml:"name"`
}

