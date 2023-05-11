package cdk8splus26


// Container lifecycle properties.
type ContainerLifecycle struct {
	// This hook is executed immediately after a container is created.
	//
	// However,
	// there is no guarantee that the hook will execute before the container ENTRYPOINT.
	PostStart Handler `field:"optional" json:"postStart" yaml:"postStart"`
	// This hook is called immediately before a container is terminated due to an API request or management event such as a liveness/startup probe failure, preemption, resource contention and others.
	//
	// A call to the PreStop hook fails if the container is already in a terminated or completed state
	// and the hook must complete before the TERM signal to stop the container can be sent.
	// The Pod's termination grace period countdown begins before the PreStop hook is executed,
	// so regardless of the outcome of the handler, the container will eventually terminate
	// within the Pod's termination grace period. No parameters are passed to the handler.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle/#pod-termination
	//
	PreStop Handler `field:"optional" json:"preStop" yaml:"preStop"`
}

