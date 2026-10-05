//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceSliceList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeResourceSliceList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceSliceList_ManifestParameters(props *KubeResourceSliceListProps) error {
	return nil
}

func validateKubeResourceSliceList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceSliceListParameters(scope constructs.Construct, id *string, props *KubeResourceSliceListProps) error {
	return nil
}

