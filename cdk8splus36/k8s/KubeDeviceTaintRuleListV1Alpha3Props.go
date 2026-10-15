package k8s


// DeviceTaintRuleList is a collection of DeviceTaintRules.
type KubeDeviceTaintRuleListV1Alpha3Props struct {
	// Items is the list of DeviceTaintRules.
	Items *[]*KubeDeviceTaintRuleV1Alpha3Props `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

