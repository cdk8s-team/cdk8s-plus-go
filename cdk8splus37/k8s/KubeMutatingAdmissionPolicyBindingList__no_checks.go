//go:build no_runtime_type_checking

package k8s

// Building without runtime type checking enabled, so all the below just return nil

func validateKubeMutatingAdmissionPolicyBindingList_IsApiObjectParameters(o interface{}) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyBindingList_IsConstructParameters(x interface{}) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyBindingList_ManifestParameters(props *KubeMutatingAdmissionPolicyBindingListProps) error {
	return nil
}

func validateKubeMutatingAdmissionPolicyBindingList_OfParameters(c constructs.IConstruct) error {
	return nil
}

func validateNewKubeMutatingAdmissionPolicyBindingListParameters(scope constructs.Construct, id *string, props *KubeMutatingAdmissionPolicyBindingListProps) error {
	return nil
}

