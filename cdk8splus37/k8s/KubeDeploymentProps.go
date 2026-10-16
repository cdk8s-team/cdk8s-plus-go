package k8s


// Deployment enables declarative updates for Pods and ReplicaSets.
type KubeDeploymentProps struct {
	// Specification of the desired behavior of the Deployment.
	Spec *DeploymentSpec `field:"required" json:"spec" yaml:"spec"`
	// Standard object's metadata.
	//
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

