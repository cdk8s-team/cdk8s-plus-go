package k8s


// CapacityRequestPolicy defines how requests consume device capacity.
//
// Must not set more than one ValidRequestValues.
type CapacityRequestPolicy struct {
	// Default specifies how much of this capacity is consumed by a request that does not contain an entry for it in DeviceRequest's Capacity.
	Default Quantity `field:"optional" json:"default" yaml:"default"`
	// ValidRange defines an acceptable quantity value range in consuming requests.
	//
	// If this field is set, Default must be defined and it must fall within the defined ValidRange.
	//
	// If the requested amount does not fall within the defined range, the request violates the policy, and this device cannot be allocated.
	//
	// If the request doesn't contain this capacity entry, Default value is used.
	ValidRange *CapacityRequestPolicyRange `field:"optional" json:"validRange" yaml:"validRange"`
	// ValidValues defines a set of acceptable quantity values in consuming requests.
	//
	// Must not contain more than 10 entries. Must be sorted in ascending order.
	//
	// If this field is set, Default must be defined and it must be included in ValidValues list.
	//
	// If the requested amount does not match any valid value but smaller than some valid values, the scheduler calculates the smallest valid value that is greater than or equal to the request. That is: min(ceil(requestedValue) ∈ validValues), where requestedValue ≤ max(validValues).
	//
	// If the requested amount exceeds all valid values, the request violates the policy, and this device cannot be allocated.
	ValidValues *[]Quantity `field:"optional" json:"validValues" yaml:"validValues"`
}

