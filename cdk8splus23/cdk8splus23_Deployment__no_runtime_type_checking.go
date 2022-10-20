//go:build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

// Building without runtime type checking enabled, so all the below just return nil

func (d *jsiiProxy_Deployment) validateAddContainerParameters(cont *ContainerProps) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateAddInitContainerParameters(cont *ContainerProps) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateAddVolumeParameters(vol Volume) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateExposeViaIngressParameters(path *string, options *ExposeDeploymentViaIngressOptions) error {
	return nil
}

func (d *jsiiProxy_Deployment) validateExposeViaServiceParameters(options *DeploymentExposeViaServiceOptions) error {
	return nil
}

func validateDeployment_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewDeploymentParameters(scope constructs.Construct, id *string, props *DeploymentProps) error {
	return nil
}

