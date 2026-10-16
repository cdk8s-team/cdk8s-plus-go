package k8s


// PodCertificateRequest encodes a pod requesting a certificate from a given signer.
//
// Kubelets use this API to implement podCertificate projected volumes.
type KubePodCertificateRequestProps struct {
	// spec contains the details about the certificate being requested.
	Spec *PodCertificateRequestSpec `field:"required" json:"spec" yaml:"spec"`
	// metadata contains the object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

