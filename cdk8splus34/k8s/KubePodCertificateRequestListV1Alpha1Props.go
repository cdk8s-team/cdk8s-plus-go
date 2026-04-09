package k8s


// PodCertificateRequestList is a collection of PodCertificateRequest objects.
type KubePodCertificateRequestListV1Alpha1Props struct {
	// items is a collection of PodCertificateRequest objects.
	Items *[]*KubePodCertificateRequestV1Alpha1Props `field:"required" json:"items" yaml:"items"`
	// metadata contains the list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

