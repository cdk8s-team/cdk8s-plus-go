package k8s


// DeviceTaintRule adds one taint to all devices which match the selector.
//
// This has the same effect as if the taint was specified directly in the ResourceSlice by the DRA driver.
type KubeDeviceTaintRuleV1Alpha3Props struct {
	// Spec specifies the selector and one taint.
	//
	// Changing the spec automatically increments the metadata.generation number.
	Spec *DeviceTaintRuleSpecV1Alpha3 `field:"required" json:"spec" yaml:"spec"`
	// Standard object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

