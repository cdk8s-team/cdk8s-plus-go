package k8s


// ContainerRestartRuleOnExitCodes describes the condition for handling an exited container based on its exit codes.
type ContainerRestartRuleOnExitCodes struct {
	// Represents the relationship between the container exit code(s) and the specified values.
	//
	// Possible values are: - In: the requirement is satisfied if the container exit code is in the
	// set of specified values.
	// - NotIn: the requirement is satisfied if the container exit code is
	// not in the set of specified values.
	Operator *string `field:"required" json:"operator" yaml:"operator"`
	// Specifies the set of values to check for container exit codes.
	//
	// At most 255 elements are allowed.
	Values *[]*float64 `field:"optional" json:"values" yaml:"values"`
}

