//go:build no_runtime_type_checking

package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func validateUser_FromNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateUser_IsConstructParameters(x interface{}) error {
	return nil
}

