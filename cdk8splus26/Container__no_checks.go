//go:build no_runtime_type_checking

package cdk8splus26

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_Container) validateAddPortParameters(port *ContainerPort) error {
	return nil
}

func (c *jsiiProxy_Container) validateMountParameters(path *string, storage IStorage, options *MountOptions) error {
	return nil
}

func validateNewContainerParameters(props *ContainerProps) error {
	return nil
}

