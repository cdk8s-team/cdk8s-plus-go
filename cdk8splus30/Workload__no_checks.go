//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_Workload) validateAddContainerParameters(cont *ContainerProps) error {
	return nil
}

func (w *jsiiProxy_Workload) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	return nil
}

func (w *jsiiProxy_Workload) validateAddInitContainerParameters(cont *ContainerProps) error {
	return nil
}

func (w *jsiiProxy_Workload) validateAddVolumeParameters(vol Volume) error {
	return nil
}

func (w *jsiiProxy_Workload) validateAttachContainerParameters(cont Container) error {
	return nil
}

func validateWorkload_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewWorkloadParameters(scope constructs.Construct, id *string, props *WorkloadProps) error {
	return nil
}

