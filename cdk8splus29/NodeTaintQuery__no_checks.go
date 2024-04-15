//go:build no_runtime_type_checking

package cdk8splus29

// Building without runtime type checking enabled, so all the below just return nil

func validateNodeTaintQuery_ExistsParameters(key *string, options *NodeTaintQueryOptions) error {
	return nil
}

func validateNodeTaintQuery_IsParameters(key *string, value *string, options *NodeTaintQueryOptions) error {
	return nil
}

