//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceSlice_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeResourceSlice_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceSlice_ManifestParameters(props *KubeResourceSliceProps) error {
	return nil
}

func validateKubeResourceSlice_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceSliceParameters(scope constructs.Construct, id *string, props *KubeResourceSliceProps) error {
	return nil
}

