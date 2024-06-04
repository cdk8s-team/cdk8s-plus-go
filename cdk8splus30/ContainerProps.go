package cdk8splus30


// Properties for creating a container.
type ContainerProps struct {
	// Arguments to the entrypoint. The docker image's CMD is used if `command` is not provided.
	//
	// Variable references $(VAR_NAME) are expanded using the container's
	// environment. If a variable cannot be resolved, the reference in the input
	// string will be unchanged. The $(VAR_NAME) syntax can be escaped with a
	// double $$, ie: $$(VAR_NAME). Escaped references will never be expanded,
	// regardless of whether the variable exists or not.
	//
	// Cannot be updated.
	// See: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	//
	// Default: [].
	//
	Args *[]*string `field:"optional" json:"args" yaml:"args"`
	// Entrypoint array.
	//
	// Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment.
	// If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME).
	// Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated.
	// More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	// Default: - The docker image's ENTRYPOINT.
	//
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// List of sources to populate environment variables in the container.
	//
	// When a key exists in multiple sources, the value associated with
	// the last source will take precedence. Values defined by the `envVariables` property
	// with a duplicate key will take precedence.
	// Default: - No sources.
	//
	EnvFrom *[]EnvFrom `field:"optional" json:"envFrom" yaml:"envFrom"`
	// Environment variables to set in the container.
	// Default: - No environment variables.
	//
	EnvVariables *map[string]EnvValue `field:"optional" json:"envVariables" yaml:"envVariables"`
	// Image pull policy for this container.
	// Default: ImagePullPolicy.ALWAYS
	//
	ImagePullPolicy ImagePullPolicy `field:"optional" json:"imagePullPolicy" yaml:"imagePullPolicy"`
	// Describes actions that the management system should take in response to container lifecycle events.
	Lifecycle *ContainerLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Periodic probe of container liveness.
	//
	// Container will be restarted if the probe fails.
	// Default: - no liveness probe is defined.
	//
	Liveness Probe `field:"optional" json:"liveness" yaml:"liveness"`
	// Name of the container specified as a DNS_LABEL.
	//
	// Each container in a pod must have a unique name (DNS_LABEL). Cannot be updated.
	// Default: 'main'.
	//
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Deprecated: - use `portNumber`.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Number of port to expose on the pod's IP address.
	//
	// This must be a valid port number, 0 < x < 65536.
	//
	// This is a convinience property if all you need a single TCP numbered port.
	// In case more advanced configuartion is required, use the `ports` property.
	//
	// This port is added to the list of ports mentioned in the `ports` property.
	// Default: - Only the ports mentiond in the `ports` property are exposed.
	//
	PortNumber *float64 `field:"optional" json:"portNumber" yaml:"portNumber"`
	// List of ports to expose from this container.
	// Default: - Only the port mentioned in the `portNumber` property is exposed.
	//
	Ports *[]*ContainerPort `field:"optional" json:"ports" yaml:"ports"`
	// Determines when the container is ready to serve traffic.
	// Default: - no readiness probe is defined.
	//
	Readiness Probe `field:"optional" json:"readiness" yaml:"readiness"`
	// Compute resources (CPU and memory requests and limits) required by the container.
	// See: https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
	//
	// Default:    cpu:
	//      request: 1000 millis
	//      limit: 1500 millis
	//    memory:
	//      request: 512 mebibytes
	// limit: 2048 mebibytes.
	//
	Resources *ContainerResources `field:"optional" json:"resources" yaml:"resources"`
	// Kubelet will start init containers with restartPolicy=Always in the order with other init containers, but instead of waiting for its completion, it will wait for the container startup completion Currently, only accepted value is Always.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/sidecar-containers/
	//
	// Default: - no restart policy is defined and the pod restart policy is applied.
	//
	RestartPolicy ContainerRestartPolicy `field:"optional" json:"restartPolicy" yaml:"restartPolicy"`
	// SecurityContext defines the security options the container should be run with.
	//
	// If set, the fields override equivalent fields of the pod's security context.
	// See: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	//
	// Default:   ensureNonRoot: true
	//   privileged: false
	//   readOnlyRootFilesystem: true
	//   allowPrivilegeEscalation: false
	//   user: 25000
	// group: 26000.
	//
	SecurityContext *ContainerSecurityContextProps `field:"optional" json:"securityContext" yaml:"securityContext"`
	// StartupProbe indicates that the Pod has successfully initialized.
	//
	// If specified, no other probes are executed until this completes successfully.
	// Default: - If a port is provided, then knocks on that port
	// to determine when the container is ready for readiness and
	// liveness probe checks.
	// Otherwise, no startup probe is defined.
	//
	Startup Probe `field:"optional" json:"startup" yaml:"startup"`
	// Pod volumes to mount into the container's filesystem.
	//
	// Cannot be updated.
	VolumeMounts *[]*VolumeMount `field:"optional" json:"volumeMounts" yaml:"volumeMounts"`
	// Container's working directory.
	//
	// If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	// Default: - The container runtime's default.
	//
	WorkingDir *string `field:"optional" json:"workingDir" yaml:"workingDir"`
	// Docker image name.
	Image *string `field:"required" json:"image" yaml:"image"`
}

