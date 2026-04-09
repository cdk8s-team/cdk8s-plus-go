//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeResourceClaimTemplateList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeResourceClaimTemplateList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeResourceClaimTemplateList_ManifestParameters(props *KubeResourceClaimTemplateListProps) error {
	return nil
}

func validateKubeResourceClaimTemplateList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeResourceClaimTemplateListParameters(scope constructs.Construct, id *string, props *KubeResourceClaimTemplateListProps) error {
	return nil
}

