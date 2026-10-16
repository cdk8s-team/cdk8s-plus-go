package k8s


// DeviceTaintRuleList is a collection of DeviceTaintRules.
type KubeDeviceTaintRuleListProps struct {
	// Items is the list of DeviceTaintRules.
	Items *[]*KubeDeviceTaintRuleProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

