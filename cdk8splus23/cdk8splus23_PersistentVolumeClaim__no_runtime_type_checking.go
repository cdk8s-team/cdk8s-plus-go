//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PersistentVolumeClaim) validateBindParameters(vol IPersistentVolume) error {
	return nil
}

func validatePersistentVolumeClaim_FromClaimNameParameters(scope constructs.Construct, id *string, claimName *string) error {
	return nil
}

func validatePersistentVolumeClaim_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPersistentVolumeClaimParameters(scope constructs.Construct, id *string, props *PersistentVolumeClaimProps) error {
	return nil
}

