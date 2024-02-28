//go:build no_runtime_type_checking

package cdk8splus28

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceAccount) validateAddSecretParameters(secr ISecret) error {
	return nil
}

func validateServiceAccount_FromServiceAccountNameParameters(scope constructs.Construct, id *string, name *string, options *FromServiceAccountNameOptions) error {
	return nil
}

func validateServiceAccount_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewServiceAccountParameters(scope constructs.Construct, id *string, props *ServiceAccountProps) error {
	return nil
}

