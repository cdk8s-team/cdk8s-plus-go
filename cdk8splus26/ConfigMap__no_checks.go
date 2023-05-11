//go:build no_runtime_type_checking

package cdk8splus26

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_ConfigMap) validateAddBinaryDataParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateAddDataParameters(key *string, value *string) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateAddDirectoryParameters(localDir *string, options *AddDirectoryOptions) error {
	return nil
}

func (c *jsiiProxy_ConfigMap) validateAddFileParameters(localFile *string) error {
	return nil
}

func validateConfigMap_FromConfigMapNameParameters(scope constructs.Construct, id *string, name *string) error {
	return nil
}

func validateConfigMap_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewConfigMapParameters(scope constructs.Construct, id *string, props *ConfigMapProps) error {
	return nil
}

