package cdk8splus30


type ConnectionScheme string

const (
	// Use HTTP request for connecting to host.
	ConnectionScheme_HTTP ConnectionScheme = "HTTP"
	// Use HTTPS request for connecting to host.
	ConnectionScheme_HTTPS ConnectionScheme = "HTTPS"
)

