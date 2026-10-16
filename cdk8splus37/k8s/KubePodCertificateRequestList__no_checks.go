//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubePodCertificateRequestList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubePodCertificateRequestList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubePodCertificateRequestList_ManifestParameters(props *KubePodCertificateRequestListProps) error {
	return nil
}

func validateKubePodCertificateRequestList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubePodCertificateRequestListParameters(scope constructs.Construct, id *string, props *KubePodCertificateRequestListProps) error {
	return nil
}

