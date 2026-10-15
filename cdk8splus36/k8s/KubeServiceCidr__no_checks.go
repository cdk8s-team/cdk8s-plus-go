//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeServiceCidr_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeServiceCidr_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeServiceCidr_ManifestParameters(props *KubeServiceCidrProps) error {
	return nil
}

func validateKubeServiceCidr_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeServiceCidrParameters(scope constructs.Construct, id *string, props *KubeServiceCidrProps) error {
	return nil
}

