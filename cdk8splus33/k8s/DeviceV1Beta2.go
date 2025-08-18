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
	// Attributes defines the set of attributes for this device.
	//
	// The name of each attribute must be unique in that set.
	//
	// The maximum number of attributes and capacities combined is 32.
	Attributes *map[string]*DeviceAttributeV1Beta2 `field:"optional" json:"attributes" yaml:"attributes"`
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
	// The total number of device counter consumption entries must be <= 32. In addition, the total number in the entire ResourceSlice must be <= 1024 (for example, 64 devices with 16 counters each).
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
	// The maximum number of taints is 4.
	//
	// This is an alpha field and requires enabling the DRADeviceTaints feature gate.
	Taints *[]*DeviceTaintV1Beta2 `field:"optional" json:"taints" yaml:"taints"`
}

