package cdk8splus26


// Properties used to configure the target of an Autoscaler.
type ScalingTarget struct {
	// The object's API version (e.g. "authorization.k8s.io/v1").
	ApiVersion *string `field:"required" json:"apiVersion" yaml:"apiVersion"`
	// Container definitions associated with the target.
	Containers *[]Container `field:"required" json:"containers" yaml:"containers"`
	// The object kind (e.g. "Deployment").
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// The Kubernetes name of this resource.
	Name *string `field:"required" json:"name" yaml:"name"`
	// The fixed number of replicas defined on the target.
	//
	// This is used
	// for validation purposes as Scalable targets should not have a
	// fixed number of replicas.
	Replicas *float64 `field:"optional" json:"replicas" yaml:"replicas"`
}

