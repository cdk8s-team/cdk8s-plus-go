//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceClaimList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeResourceClaimList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceClaimList_ManifestParameters(props *KubeResourceClaimListProps) error {
	return nil
}

func validateKubeResourceClaimList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceClaimListParameters(scope constructs.Construct, id *string, props *KubeResourceClaimListProps) error {
	return nil
}

