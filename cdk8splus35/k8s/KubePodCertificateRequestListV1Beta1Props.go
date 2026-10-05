package k8s


// PodCertificateRequestList is a collection of PodCertificateRequest objects.
type KubePodCertificateRequestListV1Beta1Props struct {
	// items is a collection of PodCertificateRequest objects.
	Items *[]*KubePodCertificateRequestV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// metadata contains the list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

