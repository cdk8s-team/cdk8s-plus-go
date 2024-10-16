//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func (s *jsiiProxy_ServiceAccountTokenSecret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (s *jsiiProxy_ServiceAccountTokenSecret) validateEnvValueParameters(key *string, options *EnvValueFromSecretOptions) error {
	return nil
}

func (s *jsiiProxy_ServiceAccountTokenSecret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateServiceAccountTokenSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateServiceAccountTokenSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewServiceAccountTokenSecretParameters(scope constructs.Construct, id *string, props *ServiceAccountTokenSecretProps) error {
	return nil
}

