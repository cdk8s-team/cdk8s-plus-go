package k8s


// What happens to a persistent volume when released from its claim.
//
// Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
//
// Possible enum values:
// - `"Delete"` means the volume will be deleted from Kubernetes on release from its claim. The volume plugin must support Deletion.
// - `"Recycle"` means the volume will be recycled back into the pool of unbound persistent volumes on release from its claim. The volume plugin must support Recycling.
// - `"Retain"` means the volume will be left in its current phase (Released) for manual reclamation by the administrator. The default policy is Retain.
type IoK8SApiCoreV1PersistentVolumeSpecPersistentVolumeReclaimPolicy string

const (
	// Delete.
	IoK8SApiCoreV1PersistentVolumeSpecPersistentVolumeReclaimPolicy_DELETE IoK8SApiCoreV1PersistentVolumeSpecPersistentVolumeReclaimPolicy = "DELETE"
	// Recycle.
	IoK8SApiCoreV1PersistentVolumeSpecPersistentVolumeReclaimPolicy_RECYCLE IoK8SApiCoreV1PersistentVolumeSpecPersistentVolumeReclaimPolicy = "RECYCLE"
	// Retain.
	IoK8SApiCoreV1PersistentVolumeSpecPersistentVolumeReclaimPolicy_RETAIN IoK8SApiCoreV1PersistentVolumeSpecPersistentVolumeReclaimPolicy = "RETAIN"
)

