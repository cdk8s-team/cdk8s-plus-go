//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_DaemonSet) validateAddContainerParameters(cont *ContainerProps) error {
	return nil
}

func (d *jsiiProxy_DaemonSet) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	return nil
}

func (d *jsiiProxy_DaemonSet) validateAddInitContainerParameters(cont *ContainerProps) error {
	return nil
}

func (d *jsiiProxy_DaemonSet) validateAddVolumeParameters(vol Volume) error {
	return nil
}

func validateDaemonSet_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDaemonSetParameters(scope constructs.Construct, id *string, props *DaemonSetProps) error {
	return nil
}

