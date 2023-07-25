//go:build no_runtime_type_checking

package cdk8splus27

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

