//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

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

