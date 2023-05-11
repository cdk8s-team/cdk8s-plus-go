//go:build no_runtime_type_checking

package cdk8splus26

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AwsElasticBlockStorePersistentVolume) validateBindParameters(claim IPersistentVolumeClaim) error {
	return nil
}

func validateAwsElasticBlockStorePersistentVolume_FromPersistentVolumeNameParameters(scope constructs.Construct, id *string, volumeName *string) error {
	return nil
}

func validateAwsElasticBlockStorePersistentVolume_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAwsElasticBlockStorePersistentVolumeParameters(scope constructs.Construct, id *string, props *AwsElasticBlockStorePersistentVolumeProps) error {
	return nil
}

