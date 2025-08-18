//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeIpAddress_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeIpAddress_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeIpAddress_ManifestParameters(props *KubeIpAddressProps) error {
	return nil
}

func validateKubeIpAddress_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeIpAddressParameters(scope constructs.Construct, id *string, props *KubeIpAddressProps) error {
	return nil
}

