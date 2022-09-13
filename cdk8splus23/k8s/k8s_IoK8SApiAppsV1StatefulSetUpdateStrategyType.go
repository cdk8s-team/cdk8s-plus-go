package k8s


// Type indicates the type of the StatefulSetUpdateStrategy. Default is RollingUpdate.
//
// Possible enum values:
// - `"OnDelete"` triggers the legacy behavior. Version tracking and ordered rolling restarts are disabled. Pods are recreated from the StatefulSetSpec when they are manually deleted. When a scale operation is performed with this strategy,specification version indicated by the StatefulSet's currentRevision.
// - `"RollingUpdate"` indicates that update will be applied to all Pods in the StatefulSet with respect to the StatefulSet ordering constraints. When a scale operation is performed with this strategy, new Pods will be created from the specification version indicated by the StatefulSet's updateRevision.
type IoK8SApiAppsV1StatefulSetUpdateStrategyType string

const (
	// OnDelete.
	IoK8SApiAppsV1StatefulSetUpdateStrategyType_ON_DELETE IoK8SApiAppsV1StatefulSetUpdateStrategyType = "ON_DELETE"
	// RollingUpdate.
	IoK8SApiAppsV1StatefulSetUpdateStrategyType_ROLLING_UPDATE IoK8SApiAppsV1StatefulSetUpdateStrategyType = "ROLLING_UPDATE"
)

