//go:build no_runtime_type_checking

package cdk8splus27

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DockerConfigSecret) validateAddStringDataParameters(key *string, value *string) error {
	return nil
}

func (d *jsiiProxy_DockerConfigSecret) validateEnvValueParameters(key *string, options *EnvValueFromSecretOptions) error {
	return nil
}

func (d *jsiiProxy_DockerConfigSecret) validateGetStringDataParameters(key *string) error {
	return nil
}

func validateDockerConfigSecret_FromSecretNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateDockerConfigSecret_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDockerConfigSecretParameters(scope constructs.Construct, id *string, props *DockerConfigSecretProps) error {
	return nil
}

