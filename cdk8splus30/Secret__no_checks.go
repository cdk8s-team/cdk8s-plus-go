//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_Secret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_Secret) validateEnvValueParameters(key *string, options *EnvValueFromSecretOptions) error {
	return nil
}

func (s *jsiiProxy_Secret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSecretParameters(scope constructs.Construct, id *string, props *SecretProps) error {
	return nil
}

