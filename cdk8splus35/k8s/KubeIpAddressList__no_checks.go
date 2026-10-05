//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeIpAddressList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeIpAddressList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeIpAddressList_ManifestParameters(props *KubeIpAddressListProps) error {
	return nil
}

func validateKubeIpAddressList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeIpAddressListParameters(scope constructs.Construct, id *string, props *KubeIpAddressListProps) error {
	return nil
}

