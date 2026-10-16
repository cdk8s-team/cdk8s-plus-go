//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeDeviceTaintRule_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeDeviceTaintRule_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeDeviceTaintRule_ManifestParameters(props *KubeDeviceTaintRuleProps) error {
	return nil
}

func validateKubeDeviceTaintRule_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeDeviceTaintRuleParameters(scope constructs.Construct, id *string, props *KubeDeviceTaintRuleProps) error {
	return nil
}

