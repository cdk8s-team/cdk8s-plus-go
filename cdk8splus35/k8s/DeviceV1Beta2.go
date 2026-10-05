package k8s


// Device represents one individual hardware instance that can be selected based on its attributes.
//
// Besides the name, exactly one field must be set.
type DeviceV1Beta2 struct {
	// Name is unique identifier among all devices managed by the driver in the pool.
	//
	// It must be a DNS label.
	Name *string `field:"required" json:"name" yaml:"name"`
	// AllNodes indicates that all nodes have access to the device.
	//
	// Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
	AllNodes *bool `field:"optional" json:"allNodes" yaml:"allNodes"`
	// AllowMultipleAllocations marks whether the device is allowed to be allocated to multiple DeviceRequests.
	//
	// If AllowMultipleAllocations is set to true, the device can be allocated more than once, and all of its capacity is consumable, regardless of whether the requestPolicy is defined or not.
	AllowMultipleAllocations *bool `field:"optional" json:"allowMultipleAllocations" yaml:"allowMultipleAllocations"`
	// Attributes defines the set of attributes for this device.
	//
	// The name of each attribute must be unique in that set.
	//
	// The maximum number of attributes and capacities combined is 32.
	Attributes *map[string]*DeviceAttributeV1Beta2 `field:"optional" json:"attributes" yaml:"attributes"`
	// BindingConditions defines the conditions for proceeding with binding.
	//
	// All of these conditions must be set in the per-device status conditions with a value of True to proceed with binding the pod to the node while scheduling the pod.
	//
	// The maximum number of binding conditions is 4.
	//
	// The conditions must be a valid condition type string.
	//
	// This is an alpha field and requires enabling the DRADeviceBindingConditions and DRAResourceClaimDeviceStatus feature gates.
	BindingConditions *[]*string `field:"optional" json:"bindingConditions" yaml:"bindingConditions"`
	// BindingFailureConditions defines the conditions for binding failure.
	//
	// They may be set in the per-device status conditions. If any is set to "True", a binding failure occurred.
	//
	// The maximum number of binding failure conditions is 4.
	//
	// The conditions must be a valid condition type string.
	//
	// This is an alpha field and requires enabling the DRADeviceBindingConditions and DRAResourceClaimDeviceStatus feature gates.
	BindingFailureConditions *[]*string `field:"optional" json:"bindingFailureConditions" yaml:"bindingFailureConditions"`
	// BindsToNode indicates if the usage of an allocation involving this device has to be limited to exactly the node that was chosen when allocating the claim.
	//
	// If set to true, the scheduler will set the ResourceClaim.Status.Allocation.NodeSelector to match the node where the allocation was made.
	//
	// This is an alpha field and requires enabling the DRADeviceBindingConditions and DRAResourceClaimDeviceStatus feature gates.
	BindsToNode *bool `field:"optional" json:"bindsToNode" yaml:"bindsToNode"`
	// Capacity defines the set of capacities for this device.
	//
	// The name of each capacity must be unique in that set.
	//
	// The maximum number of attributes and capacities combined is 32.
	Capacity *map[string]*DeviceCapacityV1Beta2 `field:"optional" json:"capacity" yaml:"capacity"`
	// ConsumesCounters defines a list of references to sharedCounters and the set of counters that the device will consume from those counter sets.
	//
	// There can only be a single entry per counterSet.
	//
	// The maximum number of device counter consumptions per device is 2.
	ConsumesCounters *[]*DeviceCounterConsumptionV1Beta2 `field:"optional" json:"consumesCounters" yaml:"consumesCounters"`
	// NodeName identifies the node where the device is available.
	//
	// Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
	NodeName *string `field:"optional" json:"nodeName" yaml:"nodeName"`
	// NodeSelector defines the nodes where the device is available.
	//
	// Must use exactly one term.
	//
	// Must only be set if Spec.PerDeviceNodeSelection is set to true. At most one of NodeName, NodeSelector and AllNodes can be set.
	NodeSelector *NodeSelector `field:"optional" json:"nodeSelector" yaml:"nodeSelector"`
	// If specified, these are the driver-defined taints.
	//
	// The maximum number of taints is 16. If taints are set for any device in a ResourceSlice, then the maximum number of allowed devices per ResourceSlice is 64 instead of 128.
	//
	// This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Taints *[]*DeviceTaintV1Beta2 `field:"optional" json:"taints" yaml:"taints"`
}

