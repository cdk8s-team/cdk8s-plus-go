package k8s


// type indicates which kind of seccomp profile will be applied. Valid options are:.
//
// Localhost - a profile defined in a file on the node should be used. RuntimeDefault - the container runtime default profile should be used. Unconfined - no profile should be applied.
//
// Possible enum values:
// - `"Localhost"` indicates a profile defined in a file on the node should be used. The file's location relative to <kubelet-root-dir>/seccomp.
// - `"RuntimeDefault"` represents the default container runtime seccomp profile.
// - `"Unconfined"` indicates no seccomp profile is applied (A.K.A. unconfined).
type IoK8SApiCoreV1SeccompProfileType string

const (
	// Localhost.
	IoK8SApiCoreV1SeccompProfileType_LOCALHOST IoK8SApiCoreV1SeccompProfileType = "LOCALHOST"
	// RuntimeDefault.
	IoK8SApiCoreV1SeccompProfileType_RUNTIME_DEFAULT IoK8SApiCoreV1SeccompProfileType = "RUNTIME_DEFAULT"
	// Unconfined.
	IoK8SApiCoreV1SeccompProfileType_UNCONFINED IoK8SApiCoreV1SeccompProfileType = "UNCONFINED"
)

