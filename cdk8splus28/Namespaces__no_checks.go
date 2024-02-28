//go:build no_runtime_type_checking

package cdk8splus28

// Building without runtime type checking enabled, so all the below just return nil

func validateNamespaces_AllParameters(scope constructs.Construct, id *string) error {
	return nil
}

func validateNamespaces_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNamespaces_SelectParameters(scope constructs.Construct, id *string, options *NamespacesSelectOptions) error {
	return nil
}

func validateNewNamespacesParameters(scope constructs.Construct, id *string) error {
	return nil
}

