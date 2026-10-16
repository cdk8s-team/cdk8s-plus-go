//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeMutatingAdmissionPolicyBinding_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyBinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyBinding_ManifestParameters(props *KubeMutatingAdmissionPolicyBindingProps) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyBinding_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeMutatingAdmissionPolicyBindingParameters(scope constructs.Construct, id *string, props *KubeMutatingAdmissionPolicyBindingProps) error {
	return nil
}

