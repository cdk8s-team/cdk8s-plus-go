//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func (j *jsiiProxy_Job) validateAddContainerParameters(cont *ContainerProps) error {
	return nil
}

func (j *jsiiProxy_Job) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	return nil
}

func (j *jsiiProxy_Job) validateAddInitContainerParameters(cont *ContainerProps) error {
	return nil
}

func (j *jsiiProxy_Job) validateAddVolumeParameters(vol Volume) error {
	return nil
}

func (j *jsiiProxy_Job) validateAttachContainerParameters(cont Container) error {
	return nil
}

func validateJob_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewJobParameters(scope constructs.Construct, id *string, props *JobProps) error {
	return nil
}

