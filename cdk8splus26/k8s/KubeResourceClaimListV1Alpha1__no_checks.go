//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceClaimListV1Alpha1_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceClaimListV1Alpha1_ManifestParameters(props *KubeResourceClaimListV1Alpha1Props) error {
	return nil
}

func validateKubeResourceClaimListV1Alpha1_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceClaimListV1Alpha1Parameters(scope constructs.Construct, id *string, props *KubeResourceClaimListV1Alpha1Props) error {
	return nil
}

