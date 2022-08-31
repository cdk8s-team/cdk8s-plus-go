//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

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

