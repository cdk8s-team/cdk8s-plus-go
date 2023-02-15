package k8s


// Scheme to use for connecting to the host. Defaults to HTTP.
//
// Possible enum values:
// - `"HTTP"` means that the scheme used will be http://
// - `"HTTPS"` means that the scheme used will be https://.
type IoK8SApiCoreV1HttpGetActionScheme string

const (
	// HTTP.
	IoK8SApiCoreV1HttpGetActionScheme_HTTP IoK8SApiCoreV1HttpGetActionScheme = "HTTP"
	// HTTPS.
	IoK8SApiCoreV1HttpGetActionScheme_HTTPS IoK8SApiCoreV1HttpGetActionScheme = "HTTPS"
)

