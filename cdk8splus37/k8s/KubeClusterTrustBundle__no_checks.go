//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeClusterTrustBundle_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeClusterTrustBundle_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeClusterTrustBundle_ManifestParameters(props *KubeClusterTrustBundleProps) error {
	return nil
}

func validateKubeClusterTrustBundle_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeClusterTrustBundleParameters(scope constructs.Construct, id *string, props *KubeClusterTrustBundleProps) error {
	return nil
}

