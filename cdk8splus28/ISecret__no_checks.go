//go:build no_runtime_type_checking

package cdk8splus28

// Building without runtime type checking enabled, so all the below just return nil

func (i *jsiiProxy_ISecret) validateEnvValueParameters(key *string, options *EnvValueFromSecretOptions) error {
	return nil
}

