//go:build no_runtime_type_checking

package cdk8splus26

// Building without runtime type checking enabled, so all the below just return nil

func validateLabelExpression_DoesNotExistParameters(key *string) error {
	return nil
}

func validateLabelExpression_ExistsParameters(key *string) error {
	return nil
}

func validateLabelExpression_InParameters(key *string, values *[]*string) error {
	return nil
}

func validateLabelExpression_NotInParameters(key *string, values *[]*string) error {
	return nil
}

