//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceClaim_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeResourceClaim_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceClaim_ManifestParameters(props *KubeResourceClaimProps) error {
	return nil
}

func validateKubeResourceClaim_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceClaimParameters(scope constructs.Construct, id *string, props *KubeResourceClaimProps) error {
	return nil
}

