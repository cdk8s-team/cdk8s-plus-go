//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceClaimTemplate_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeResourceClaimTemplate_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceClaimTemplate_ManifestParameters(props *KubeResourceClaimTemplateProps) error {
	return nil
}

func validateKubeResourceClaimTemplate_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceClaimTemplateParameters(scope constructs.Construct, id *string, props *KubeResourceClaimTemplateProps) error {
	return nil
}

