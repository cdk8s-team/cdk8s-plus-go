// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Restart policy for all containers within the pod.
type RestartPolicy string

const (
	// Always restart the pod after it exits.
	RestartPolicy_ALWAYS RestartPolicy = "ALWAYS"
	// Only restart if the pod exits with a non-zero exit code.
	RestartPolicy_ON_FAILURE RestartPolicy = "ON_FAILURE"
	// Never restart the pod.
	RestartPolicy_NEVER RestartPolicy = "NEVER"
)

