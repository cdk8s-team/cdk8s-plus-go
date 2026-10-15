package k8s


// ResourcePoolStatusRequest triggers a one-time calculation of resource pool status based on the provided filters.
//
// Once status is set, the request is considered complete and will not be reprocessed. Users should delete and recreate requests to get updated information.
type KubeResourcePoolStatusRequestV1Alpha3Props struct {
	// Standard object metadata.
	Metadata *ObjectMeta `field:"required" json:"metadata" yaml:"metadata"`
	// Spec defines the filters for which pools to include in the status.
	//
	// The spec is immutable once created.
	Spec *ResourcePoolStatusRequestSpecV1Alpha3 `field:"required" json:"spec" yaml:"spec"`
}

