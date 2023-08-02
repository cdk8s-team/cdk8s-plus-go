package cdk8splus26


// Represents the TLS configuration mapping that is passed to the ingress controller for SSL termination.
type IngressTls struct {
	// Hosts are a list of hosts included in the TLS certificate.
	//
	// The values in
	// this list must match the name/s used in the TLS Secret.
	// Default: - If unspecified, it defaults to the wildcard host setting for
	// the loadbalancer controller fulfilling this Ingress.
	//
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Secret is the secret that contains the certificate and key used to terminate SSL traffic on 443.
	//
	// If the SNI host in a listener conflicts with
	// the "Host" header field used by an IngressRule, the SNI host is used for
	// termination and value of the Host header is used for routing.
	// Default: - If unspecified, it allows SSL routing based on SNI hostname.
	//
	Secret ISecret `field:"optional" json:"secret" yaml:"secret"`
}

