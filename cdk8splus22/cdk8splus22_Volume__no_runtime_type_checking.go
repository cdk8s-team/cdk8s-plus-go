//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

// Building without runtime type checking enabled, so all the below just return nil

func validateVolume_FromAwsElasticBlockStoreParameters(scope constructs.Construct, id *string, volumeId *string, options *AwsElasticBlockStoreVolumeOptions) error {
	return nil
}

func validateVolume_FromAzureDiskParameters(scope constructs.Construct, id *string, diskName *string, diskUri *string, options *AzureDiskVolumeOptions) error {
	return nil
}

func validateVolume_FromConfigMapParameters(scope constructs.Construct, id *string, configMap IConfigMap, options *ConfigMapVolumeOptions) error {
	return nil
}

func validateVolume_FromEmptyDirParameters(scope constructs.Construct, id *string, name *string, options *EmptyDirVolumeOptions) error {
	return nil
}

func validateVolume_FromGcePersistentDiskParameters(scope constructs.Construct, id *string, pdName *string, options *GCEPersistentDiskVolumeOptions) error {
	return nil
}

func validateVolume_FromHostPathParameters(scope constructs.Construct, id *string, name *string, options *HostPathVolumeOptions) error {
	return nil
}

func validateVolume_FromPersistentVolumeClaimParameters(scope constructs.Construct, id *string, claim IPersistentVolumeClaim, options *PersistentVolumeClaimVolumeOptions) error {
	return nil
}

func validateVolume_FromSecretParameters(scope constructs.Construct, id *string, secr ISecret, options *SecretVolumeOptions) error {
	return nil
}

func validateVolume_IsConstructParameters(x interface{}) error {
	return nil
}

