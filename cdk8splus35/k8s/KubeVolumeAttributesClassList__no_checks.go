//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeVolumeAttributesClassList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeVolumeAttributesClassList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeVolumeAttributesClassList_ManifestParameters(props *KubeVolumeAttributesClassListProps) error {
	return nil
}

func validateKubeVolumeAttributesClassList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeVolumeAttributesClassListParameters(scope constructs.Construct, id *string, props *KubeVolumeAttributesClassListProps) error {
	return nil
}

