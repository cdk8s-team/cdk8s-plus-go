//go:build no_runtime_type_checking
// +build no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-25 synthesizes Kubernetes manifests for Kubernetes 1.25.0
package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func (p *jsiiProxy_PodScheduling) validateAssignParameters(node NamedNode) error {
	return nil
}

func (p *jsiiProxy_PodScheduling) validateAttractParameters(node LabeledNode, options *PodSchedulingAttractOptions) error {
	return nil
}

func (p *jsiiProxy_PodScheduling) validateColocateParameters(selector IPodSelector, options *PodSchedulingColocateOptions) error {
	return nil
}

func (p *jsiiProxy_PodScheduling) validateSeparateParameters(selector IPodSelector, options *PodSchedulingSeparateOptions) error {
	return nil
}

func (p *jsiiProxy_PodScheduling) validateTolerateParameters(node TaintedNode) error {
	return nil
}

func validateNewPodSchedulingParameters(instance AbstractPod) error {
	return nil
}

