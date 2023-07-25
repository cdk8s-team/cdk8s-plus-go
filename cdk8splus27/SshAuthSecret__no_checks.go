//go:build no_runtime_type_checking

package cdk8splus27

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_SshAuthSecret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_SshAuthSecret) validateEnvValueParameters(key *string, options *EnvValueFromSecretOptions) error {
	return nil
}

func (s *jsiiProxy_SshAuthSecret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateSshAuthSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateSshAuthSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewSshAuthSecretParameters(scope constructs.Construct, id *string, props *SshAuthSecretProps) error {
	return nil
}

