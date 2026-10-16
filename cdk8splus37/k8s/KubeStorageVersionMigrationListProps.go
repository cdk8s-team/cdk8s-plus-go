package k8s


// StorageVersionMigrationList is a collection of storage version migrations.
type KubeStorageVersionMigrationListProps struct {
	// Items is the list of StorageVersionMigration.
	Items *[]*KubeStorageVersionMigrationProps `field:"required" json:"items" yaml:"items"`
	// Standard list metadata More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata.
	Metadata *ListMeta `field:"optional" json:"metadata" yaml:"metadata"`
}

