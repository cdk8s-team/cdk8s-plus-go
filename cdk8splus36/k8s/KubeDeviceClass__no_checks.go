//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeDeviceClass_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeDeviceClass_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeDeviceClass_ManifestParameters(props *KubeDeviceClassProps) error {
	return nil
}

func validateKubeDeviceClass_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeDeviceClassParameters(scope constructs.Construct, id *string, props *KubeDeviceClassProps) error {
	return nil
}

