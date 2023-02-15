package k8s


// DaemonSetUpdateStrategy is a struct used to control the update strategy for a DaemonSet.
type DaemonSetUpdateStrategy struct {
	// Rolling update config params.
	//
	// Present only if type = "RollingUpdate".
	RollingUpdate *RollingUpdateDaemonSet `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
	//
	// Possible enum values:
	// - `"OnDelete"` Replace the old daemons only when it's killed
	// - `"RollingUpdate"` Replace the old daemons by new ones using rolling update i.e replace them on each node one after the other.
	Type IoK8SApiAppsV1DaemonSetUpdateStrategyType `field:"optional" json:"type" yaml:"type"`
}

