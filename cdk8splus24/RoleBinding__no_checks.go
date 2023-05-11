//go:build no_runtime_type_checking

package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func validateRoleBinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRoleBindingParameters(scope constructs.Construct, id *string, props *RoleBindingProps) error {
	return nil
}

