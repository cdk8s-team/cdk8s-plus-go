//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AzureDiskPersistentVolume) validateBindParameters(claim IPersistentVolumeClaim) error {
	return nil
}

func validateAzureDiskPersistentVolume_FromPersistentVolumeNameParameters(scope constructs.Construct, id *string, volumeName *string) error {
	return nil
}

func validateAzureDiskPersistentVolume_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAzureDiskPersistentVolumeParameters(scope constructs.Construct, id *string, props *AzureDiskPersistentVolumeProps) error {
	return nil
}

