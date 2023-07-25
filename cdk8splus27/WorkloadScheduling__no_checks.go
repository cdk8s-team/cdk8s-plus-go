//go:build no_runtime_type_checking

package cdk8splus27

// Building without runtime type checking enabled, so all the below just return nil

func (w *jsiiProxy_WorkloadScheduling) validateAssignParameters(node NamedNode) error {
	return nil
}

func (w *jsiiProxy_WorkloadScheduling) validateAttractParameters(node LabeledNode, options *PodSchedulingAttractOptions) error {
	return nil
}

func (w *jsiiProxy_WorkloadScheduling) validateColocateParameters(selector IPodSelector, options *PodSchedulingColocateOptions) error {
	return nil
}

func (w *jsiiProxy_WorkloadScheduling) validateSeparateParameters(selector IPodSelector, options *PodSchedulingSeparateOptions) error {
	return nil
}

func (w *jsiiProxy_WorkloadScheduling) validateSpreadParameters(options *WorkloadSchedulingSpreadOptions) error {
	return nil
}

func (w *jsiiProxy_WorkloadScheduling) validateTolerateParameters(node TaintedNode) error {
	return nil
}

func validateNewWorkloadSchedulingParameters(instance AbstractPod) error {
	return nil
}

