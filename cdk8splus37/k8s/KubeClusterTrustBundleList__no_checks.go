//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeClusterTrustBundleList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeClusterTrustBundleList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeClusterTrustBundleList_ManifestParameters(props *KubeClusterTrustBundleListProps) error {
	return nil
}

func validateKubeClusterTrustBundleList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeClusterTrustBundleListParameters(scope constructs.Construct, id *string, props *KubeClusterTrustBundleListProps) error {
	return nil
}

