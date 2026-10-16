package k8s


// ClusterTrustBundleList is a collection of ClusterTrustBundle objects.
type KubeClusterTrustBundleListProps struct {
	// items is a collection of ClusterTrustBundle objects.
	Items *[]*KubeClusterTrustBundleProps `field:"required" json:"items" yaml:"items"`
	// metadata contains the list metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

