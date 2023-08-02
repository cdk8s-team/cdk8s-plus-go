package cdk8splus27

import (
	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Options for `Probe.fromCommand()`.
type CommandProbeOptions struct {
	// Minimum consecutive failures for the probe to be considered failed after having succeeded.
	//
	// Defaults to 3. Minimum value is 1.
	// Default: 3.
	//
	FailureThreshold *float64 `field:"optional" json:"failureThreshold" yaml:"failureThreshold"`
	// Number of seconds after the container has started before liveness probes are initiated.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	// Default: - immediate.
	//
	InitialDelaySeconds cdk8s.Duration `field:"optional" json:"initialDelaySeconds" yaml:"initialDelaySeconds"`
	// How often (in seconds) to perform the probe.
	//
	// Default to 10 seconds. Minimum value is 1.
	// Default: Duration.seconds(10) Minimum value is 1.
	//
	PeriodSeconds cdk8s.Duration `field:"optional" json:"periodSeconds" yaml:"periodSeconds"`
	// Minimum consecutive successes for the probe to be considered successful after having failed. Defaults to 1.
	//
	// Must be 1 for liveness and startup. Minimum value is 1.
	// Default: 1 Must be 1 for liveness and startup. Minimum value is 1.
	//
	SuccessThreshold *float64 `field:"optional" json:"successThreshold" yaml:"successThreshold"`
	// Number of seconds after which the probe times out.
	//
	// Defaults to 1 second. Minimum value is 1.
	// See: https://kubernetes.io/docs/concepts/workloads/pods/pod-lifecycle#container-probes
	//
	// Default: Duration.seconds(1)
	//
	TimeoutSeconds cdk8s.Duration `field:"optional" json:"timeoutSeconds" yaml:"timeoutSeconds"`
}

