//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func (c *jsiiProxy_CronJob) validateAddContainerParameters(cont *ContainerProps) error {
	return nil
}

func (c *jsiiProxy_CronJob) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	return nil
}

func (c *jsiiProxy_CronJob) validateAddInitContainerParameters(cont *ContainerProps) error {
	return nil
}

func (c *jsiiProxy_CronJob) validateAddVolumeParameters(vol Volume) error {
	return nil
}

func (c *jsiiProxy_CronJob) validateAttachContainerParameters(cont Container) error {
	return nil
}

func validateCronJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewCronJobParameters(scope constructs.Construct, id *string, props *CronJobProps) error {
	return nil
}

