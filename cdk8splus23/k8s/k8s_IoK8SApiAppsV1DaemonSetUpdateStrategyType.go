package k8s


// Type of daemon set update. Can be "RollingUpdate" or "OnDelete". Default is RollingUpdate.
//
// Possible enum values:
// - `"OnDelete"` Replace the old daemons only when it's killed
// - `"RollingUpdate"` Replace the old daemons by new ones using rolling update i.e replace them on each node one after the other.
type IoK8SApiAppsV1DaemonSetUpdateStrategyType string

const (
	// OnDelete.
	IoK8SApiAppsV1DaemonSetUpdateStrategyType_ON_DELETE IoK8SApiAppsV1DaemonSetUpdateStrategyType = "ON_DELETE"
	// RollingUpdate.
	IoK8SApiAppsV1DaemonSetUpdateStrategyType_ROLLING_UPDATE IoK8SApiAppsV1DaemonSetUpdateStrategyType = "ROLLING_UPDATE"
)

