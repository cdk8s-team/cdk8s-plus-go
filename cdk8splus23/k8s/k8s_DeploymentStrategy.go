package k8s


// DeploymentStrategy describes how to replace existing pods with new ones.
type DeploymentStrategy struct {
	// Rolling update config params.
	//
	// Present only if DeploymentStrategyType = RollingUpdate.
	RollingUpdate *RollingUpdateDeployment `field:"optional" json:"rollingUpdate" yaml:"rollingUpdate"`
	// Type of deployment. Can be "Recreate" or "RollingUpdate". Default is RollingUpdate.
	//
	// Possible enum values:
	// - `"Recreate"` Kill all existing pods before creating new ones.
	// - `"RollingUpdate"` Replace the old ReplicaSets by new one using rolling update i.e gradually scale down the old ReplicaSets and scale up the new one.
	Type IoK8SApiAppsV1DeploymentStrategyType `field:"optional" json:"type" yaml:"type"`
}

