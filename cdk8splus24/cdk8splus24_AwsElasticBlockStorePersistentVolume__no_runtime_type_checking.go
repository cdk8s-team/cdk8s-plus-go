//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

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

