//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeNetworkPolicyList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeNetworkPolicyList_ManifestParameters(props *KubeNetworkPolicyListProps) error {
	return nil
}

func validateKubeNetworkPolicyList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeNetworkPolicyListParameters(scope constructs.Construct, id *string, props *KubeNetworkPolicyListProps) error {
	return nil
}

