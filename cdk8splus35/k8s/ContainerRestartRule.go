package k8s


// ContainerRestartRule describes how a container exit is handled.
type ContainerRestartRule struct {
	// Specifies the action taken on a container exit if the requirements are satisfied.
	//
	// The only possible value is "Restart" to restart the container.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Represents the exit codes to check on container exits.
	ExitCodes *ContainerRestartRuleOnExitCodes `field:"optional" json:"exitCodes" yaml:"exitCodes"`
}

