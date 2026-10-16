//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeMutatingAdmissionPolicyList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyList_ManifestParameters(props *KubeMutatingAdmissionPolicyListProps) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeMutatingAdmissionPolicyListParameters(scope constructs.Construct, id *string, props *KubeMutatingAdmissionPolicyListProps) error {
	return nil
}

