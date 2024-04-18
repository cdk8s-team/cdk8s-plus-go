package cdk8splus29


// RestartPolicy defines the restart behavior of individual containers in a pod.
//
// This field may only be set for init containers, and the only allowed value is "Always".
// For non-init containers or when this field is not specified,
// the restart behavior is defined by the Pod's restart policy and the container type.
// Setting the RestartPolicy as "Always" for the init container will have the following effect:
// this init container will be continually restarted on exit until all regular containers have terminated.
// Once all regular containers have completed, all init containers with restartPolicy "Always" will be shut down.
// This lifecycle differs from normal init containers and is often referred to as a "sidecar" container.
// See: https://kubernetes.io/docs/concepts/workloads/pods/sidecar-containers/
//
type ContainerRestartPolicy string

const (
	// If an init container is created with its restartPolicy set to Always, it will start and remain running during the entire life of the Pod.
	//
	// For regular containers, this is ignored by Kubernetes.
	ContainerRestartPolicy_ALWAYS ContainerRestartPolicy = "ALWAYS"
)

