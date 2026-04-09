//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeDeviceClassList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeDeviceClassList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeDeviceClassList_ManifestParameters(props *KubeDeviceClassListProps) error {
	return nil
}

func validateKubeDeviceClassList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeDeviceClassListParameters(scope constructs.Construct, id *string, props *KubeDeviceClassListProps) error {
	return nil
}

