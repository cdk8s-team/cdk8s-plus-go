package k8s


// ClusterTrustBundle is a cluster-scoped container for X.509 trust anchors (root certificates).
//
// ClusterTrustBundle objects are considered to be readable by any authenticated user in the cluster, because they can be mounted by pods using the `clusterTrustBundle` projection.  All service accounts have read access to ClusterTrustBundles by default.  Users who only have namespace-level access to a cluster can read ClusterTrustBundles by impersonating a serviceaccount that they have access to.
//
// It can be optionally associated with a particular assigner, in which case it contains one valid set of trust anchors for that signer. Signers may have multiple associated ClusterTrustBundles; each is an independent set of trust anchors for that signer. Admission control is used to enforce that only users with permissions on the signer can create or modify the corresponding bundle.
type KubeClusterTrustBundleV1Alpha1Props struct {
	// spec contains the signer (if any) and trust anchors.
	Spec *ClusterTrustBundleSpecV1Alpha1 `field:"required" json:"spec" yaml:"spec"`
	// metadata contains the object metadata.
	Metadata *ObjectMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

