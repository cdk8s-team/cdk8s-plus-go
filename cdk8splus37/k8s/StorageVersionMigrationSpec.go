package k8s


// Spec of the storage version migration.
type StorageVersionMigrationSpec struct {
	// The resource that is being migrated.
	//
	// The migrator sends requests to the endpoint serving the resource. Immutable.
	Resource *GroupResource `field:"required" json:"resource" yaml:"resource"`
}

