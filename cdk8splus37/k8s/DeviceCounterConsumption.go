package k8s


// DeviceCounterConsumption defines a set of counters that a device will consume from a CounterSet.
type DeviceCounterConsumption struct {
	// Counters defines the counters that will be consumed by the device.
	//
	// The maximum number of counters is 32.
	Counters *map[string]*Counter `field:"required" json:"counters" yaml:"counters"`
	// CounterSet is the name of the set from which the counters defined will be consumed.
	CounterSet *string `field:"required" json:"counterSet" yaml:"counterSet"`
	// CompatibilityGroups is a list of opaque group names for this counter set consumption.
	//
	// Devices that consume counters from the same counter set may only be allocated at the same time ("co-allocated") if they all share at least one common group: the intersection of the CompatibilityGroups of all co-allocated devices on that counter set must be non-empty. Devices that consume from different counter sets are never compared via this field.
	//
	// An unset field, an explicit nil, and an empty list are equivalent and mean "no groups": such a device is only co-allocatable with sibling devices on the same counter set that also have no groups, and is never co-allocatable with a device that declares one or more groups.
	//
	// Group names are opaque and meaningful only within the publishing driver's pool.
	//
	// The maximum number of groups is 2, and the names must be unique.
	CompatibilityGroups *[]*string `field:"optional" json:"compatibilityGroups" yaml:"compatibilityGroups"`
}

