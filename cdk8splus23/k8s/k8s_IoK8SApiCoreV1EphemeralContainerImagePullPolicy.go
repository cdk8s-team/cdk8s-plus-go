package k8s


// Image pull policy.
//
// One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
//
// Possible enum values:
// - `"Always"` means that kubelet always attempts to pull the latest image. Container will fail If the pull fails.
// - `"IfNotPresent"` means that kubelet pulls if the image isn't present on disk. Container will fail if the image isn't present and the pull fails.
// - `"Never"` means that kubelet never pulls an image, but only uses a local image. Container will fail if the image isn't present
type IoK8SApiCoreV1EphemeralContainerImagePullPolicy string

const (
	// Always.
	IoK8SApiCoreV1EphemeralContainerImagePullPolicy_ALWAYS IoK8SApiCoreV1EphemeralContainerImagePullPolicy = "ALWAYS"
	// IfNotPresent.
	IoK8SApiCoreV1EphemeralContainerImagePullPolicy_IF_NOT_PRESENT IoK8SApiCoreV1EphemeralContainerImagePullPolicy = "IF_NOT_PRESENT"
	// Never.
	IoK8SApiCoreV1EphemeralContainerImagePullPolicy_NEVER IoK8SApiCoreV1EphemeralContainerImagePullPolicy = "NEVER"
)

