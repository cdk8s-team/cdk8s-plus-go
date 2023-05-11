package cdk8splus26


type ImagePullPolicy string

const (
	// Every time the kubelet launches a container, the kubelet queries the container image registry to resolve the name to an image digest.
	//
	// If the kubelet has a container image with that exact
	// digest cached locally, the kubelet uses its cached image; otherwise, the kubelet downloads
	// (pulls) the image with the resolved digest, and uses that image to launch the container.
	//
	// Default is Always if ImagePullPolicy is omitted and either the image tag is :latest or
	// the image tag is omitted.
	ImagePullPolicy_ALWAYS ImagePullPolicy = "ALWAYS"
	// The image is pulled only if it is not already present locally.
	//
	// Default is IfNotPresent if ImagePullPolicy is omitted and the image tag is present but
	// not :latest.
	ImagePullPolicy_IF_NOT_PRESENT ImagePullPolicy = "IF_NOT_PRESENT"
	// The image is assumed to exist locally.
	//
	// No attempt is made to pull the image.
	ImagePullPolicy_NEVER ImagePullPolicy = "NEVER"
)

