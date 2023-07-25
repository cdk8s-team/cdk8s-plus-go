//go:build no_runtime_type_checking

package cdk8splus27

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PersistentVolume) validateBindParameters(claim IPersistentVolumeClaim) error {
	return nil
}

func validatePersistentVolume_FromPersistentVolumeNameParameters(scope constructs.Construct, id *string, volumeName *string) error {
	return nil
}

func validatePersistentVolume_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPersistentVolumeParameters(scope constructs.Construct, id *string, props *PersistentVolumeProps) error {
	return nil
}

