//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeStorageVersionMigrationList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeStorageVersionMigrationList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeStorageVersionMigrationList_ManifestParameters(props *KubeStorageVersionMigrationListProps) error {
	return nil
}

func validateKubeStorageVersionMigrationList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeStorageVersionMigrationListParameters(scope constructs.Construct, id *string, props *KubeStorageVersionMigrationListProps) error {
	return nil
}

