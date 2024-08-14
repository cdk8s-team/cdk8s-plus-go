package cdk8splus29


type SeccompProfileType string

const (
	// A profile defined in a file on the node should be used.
	SeccompProfileType_LOCALHOST SeccompProfileType = "LOCALHOST"
	// The container runtime default profile should be used.
	SeccompProfileType_RUNTIME_DEFAULT SeccompProfileType = "RUNTIME_DEFAULT"
	// No profile should be applied.
	SeccompProfileType_UNCONFINED SeccompProfileType = "UNCONFINED"
)

