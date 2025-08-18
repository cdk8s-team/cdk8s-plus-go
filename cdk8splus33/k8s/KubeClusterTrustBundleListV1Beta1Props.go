package k8s


// ClusterTrustBundleList is a collection of ClusterTrustBundle objects.
type KubeClusterTrustBundleListV1Beta1Props struct {
	// items is a collection of ClusterTrustBundle objects.
	Items *[]*KubeClusterTrustBundleV1Beta1Props `field:"required" json:"items" yaml:"items"`
	// metadata contains the list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

