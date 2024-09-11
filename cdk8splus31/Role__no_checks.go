//go:build no_runtime_type_checking

package cdk8splus31

// Building without runtime type checking enabled, so all the below just return nil

func (r *jsiiProxy_Role) validateAllowParameters(verbs *[]*string) error {
	return nil
}

func validateRole_FromRoleNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateRole_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewRoleParameters(scope constructs.Construct, id *string, props *RoleProps) error {
	return nil
}

