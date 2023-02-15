package k8s


// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
//
// Possible enum values:
// - `"Recreate"` Kill all existing pods before creating new ones.
// - `"RollingUpdate"` Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.
type IoK8SApiAppsV1DeploymentStrategyType string

const (
	// Recreate.
	IoK8SApiAppsV1DeploymentStrategyType_RECREATE IoK8SApiAppsV1DeploymentStrategyType = "RECREATE"
	// RollingUpdate.
	IoK8SApiAppsV1DeploymentStrategyType_ROLLING_UPDATE IoK8SApiAppsV1DeploymentStrategyType = "ROLLING_UPDATE"
)

