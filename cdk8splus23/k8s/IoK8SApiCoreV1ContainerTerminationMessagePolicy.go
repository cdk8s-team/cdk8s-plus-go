package k8s


// Indicate how the termination message should be populated.
//
// File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
//
// Possible enum values:
// - `"FallbackToLogsOnError"` will read the most recent contents of the container logs for the container status message when the container exits with an error and the terminationMessagePath has no contents.
// - `"File"` is the default behavior and will set the container status message to the contents of the container's terminationMessagePath when the container exits.
type IoK8SApiCoreV1ContainerTerminationMessagePolicy string

const (
	// FallbackToLogsOnError.
	IoK8SApiCoreV1ContainerTerminationMessagePolicy_FALLBACK_TO_LOGS_ON_ERROR IoK8SApiCoreV1ContainerTerminationMessagePolicy = "FALLBACK_TO_LOGS_ON_ERROR"
	// File.
	IoK8SApiCoreV1ContainerTerminationMessagePolicy_FILE IoK8SApiCoreV1ContainerTerminationMessagePolicy = "FILE"
)

