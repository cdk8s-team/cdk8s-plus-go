// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24


// Represents the TLS configuration mapping that is passed to the ingress controller for SSL termination.
type IngressTls struct {
	// Hosts are a list of hosts included in the TLS certificate.
	//
	// The values in
	// this list must match the name/s used in the TLS Secret.
	Hosts *[]*string `field:"optional" json:"hosts" yaml:"hosts"`
	// Secret is the secret that contains the certificate and key used to terminate SSL traffic on 443.
	//
	// If the SNI host in a listener conflicts with
	// the "Host" header field used by an IngressRule, the SNI host is used for
	// termination and value of the Host header is used for routing.
	Secret ISecret `field:"optional" json:"secret" yaml:"secret"`
}

