//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeMutatingAdmissionPolicy_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeMutatingAdmissionPolicy_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeMutatingAdmissionPolicy_ManifestParameters(props *KubeMutatingAdmissionPolicyProps) error {
	return nil
}

func validateKubeMutatingAdmissionPolicy_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeMutatingAdmissionPolicyParameters(scope constructs.Construct, id *string, props *KubeMutatingAdmissionPolicyProps) error {
	return nil
}

