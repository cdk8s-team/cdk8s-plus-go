//go:build no_runtime_type_checking

package cdk8splus31

// Building without runtime type checking enabled, so all the below just return nil

func validateClusterRoleBinding_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewClusterRoleBindingParameters(scope constructs.Construct, id *string, props *ClusterRoleBindingProps) error {
	return nil
}

