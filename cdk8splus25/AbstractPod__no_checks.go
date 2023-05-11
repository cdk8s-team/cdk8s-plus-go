//go:build no_runtime_type_checking

package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func (a *jsiiProxy_AbstractPod) validateAddContainerParameters(cont *ContainerProps) error {
	return nil
}

func (a *jsiiProxy_AbstractPod) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	return nil
}

func (a *jsiiProxy_AbstractPod) validateAddInitContainerParameters(cont *ContainerProps) error {
	return nil
}

func (a *jsiiProxy_AbstractPod) validateAddVolumeParameters(vol Volume) error {
	return nil
}

func (a *jsiiProxy_AbstractPod) validateAttachContainerParameters(cont Container) error {
	return nil
}

func validateAbstractPod_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewAbstractPodParameters(scope constructs.Construct, id *string, props *AbstractPodProps) error {
	return nil
}

