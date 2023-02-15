package k8s


// Type of resource that this limit applies to.
//
// Possible enum values:
// - `"Container"` Limit that applies to all containers in a namespace
// - `"PersistentVolumeClaim"` Limit that applies to all persistent volume claims in a namespace
// - `"Pod"` Limit that applies to all pods in a namespace.
type IoK8SApiCoreV1LimitRangeItemType string

const (
	// Container.
	IoK8SApiCoreV1LimitRangeItemType_CONTAINER IoK8SApiCoreV1LimitRangeItemType = "CONTAINER"
	// PersistentVolumeClaim.
	IoK8SApiCoreV1LimitRangeItemType_PERSISTENT_VOLUME_CLAIM IoK8SApiCoreV1LimitRangeItemType = "PERSISTENT_VOLUME_CLAIM"
	// Pod.
	IoK8SApiCoreV1LimitRangeItemType_POD IoK8SApiCoreV1LimitRangeItemType = "POD"
)

