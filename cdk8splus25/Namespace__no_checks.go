//go:build no_runtime_type_checking

package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func validateNamespace_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewNamespaceParameters(scope constructs.Construct, id *string, props *NamespaceProps) error {
	return nil
}

