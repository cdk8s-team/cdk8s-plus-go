//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

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

