//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_Pod) validateAddContainerParameters(cont *ContainerProps) error {
	return nil
}

func (p *jsiiProxy_Pod) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	return nil
}

func (p *jsiiProxy_Pod) validateAddInitContainerParameters(cont *ContainerProps) error {
	return nil
}

func (p *jsiiProxy_Pod) validateAddVolumeParameters(vol Volume) error {
	return nil
}

func validatePod_IsConstructParameters(x interface{}) error {
	return nil
}

func validateNewPodParameters(scope constructs.Construct, id *string, props *PodProps) error {
	return nil
}

