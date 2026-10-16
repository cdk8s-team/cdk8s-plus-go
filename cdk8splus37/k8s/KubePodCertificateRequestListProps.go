package k8s


// PodCertificateRequestList is a collection of PodCertificateRequest objects.
type KubePodCertificateRequestListProps struct {
	// items is a collection of PodCertificateRequest objects.
	Items *[]*KubePodCertificateRequestProps `field:"required" json:"items" yaml:"items"`
	// metadata contains the list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

