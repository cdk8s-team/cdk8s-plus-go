//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func (t *jsiiProxy_TlsSecret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (t *jsiiProxy_TlsSecret) validateEnvValueParameters(key *string, options *EnvValueFromSecretOptions) error {
	return nil
}

func (t *jsiiProxy_TlsSecret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateTlsSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateTlsSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewTlsSecretParameters(scope constructs.Construct, id *string, props *TlsSecretProps) error {
	return nil
}

