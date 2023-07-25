//go:build no_runtime_type_checking

package cdk8splus27

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ClusterRole) validateAggregateParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_ClusterRole) validateAllowParameters(verbs *[]*string) error {
	return nil
}

func (c *jsiiProxy_ClusterRole) validateBindInNamespaceParameters(namespace *string) error {
	return nil
}

func (c *jsiiProxy_ClusterRole) validateCombineParameters(rol ClusterRole) error {
	return nil
}

func validateClusterRole_FromClusterRoleNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateClusterRole_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewClusterRoleParameters(scope constructs.Construct, id *string, props *ClusterRoleProps) error {
	return nil
}

