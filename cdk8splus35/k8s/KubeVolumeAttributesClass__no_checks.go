//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeVolumeAttributesClass_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeVolumeAttributesClass_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeVolumeAttributesClass_ManifestParameters(props *KubeVolumeAttributesClassProps) error {
	return nil
}

func validateKubeVolumeAttributesClass_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeVolumeAttributesClassParameters(scope constructs.Construct, id *string, props *KubeVolumeAttributesClassProps) error {
	return nil
}

