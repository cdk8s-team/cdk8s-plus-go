package cdk8splus29


// Reclaim Policies.
type PersistentVolumeReclaimPolicy string

const (
	// The Retain reclaim policy allows for manual reclamation of the resource.
	//
	// When the PersistentVolumeClaim is deleted, the PersistentVolume still exists and the
	// volume is considered "released". But it is not yet available for another claim
	// because the previous claimant's data remains on the volume.
	// An administrator can manually reclaim the volume with the following steps:
	//
	// 1. Delete the PersistentVolume. The associated storage asset in external
	//     infrastructure (such as an AWS EBS, GCE PD, Azure Disk, or Cinder volume) still exists after the PV is deleted.
	// 2. Manually clean up the data on the associated storage asset accordingly.
	// 3. Manually delete the associated storage asset.
	//
	// If you want to reuse the same storage asset, create a new PersistentVolume
	// with the same storage asset definition.
	PersistentVolumeReclaimPolicy_RETAIN PersistentVolumeReclaimPolicy = "RETAIN"
	// For volume plugins that support the Delete reclaim policy, deletion removes both the PersistentVolume object from Kubernetes, as well as the associated storage asset in the external infrastructure, such as an AWS EBS, GCE PD, Azure Disk, or Cinder volume.
	//
	// Volumes that were dynamically provisioned inherit the reclaim policy of their StorageClass, which defaults to Delete.
	// The administrator should configure the StorageClass according to users' expectations; otherwise,
	// the PV must be edited or patched after it is created.
	PersistentVolumeReclaimPolicy_DELETE PersistentVolumeReclaimPolicy = "DELETE"
)

