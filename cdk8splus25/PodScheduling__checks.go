//go:build !no_runtime_type_checking

package cdk8splus25

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (p *jsiiProxy_PodScheduling) validateAssignParameters(node NamedNode) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func (p *jsiiProxy_PodScheduling) validateAttractParameters(node LabeledNode, options *PodSchedulingAttractOptions) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PodScheduling) validateColocateParameters(selector IPodSelector, options *PodSchedulingColocateOptions) error {
	if selector == nil {
		return fmt.Errorf("parameter selector is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PodScheduling) validateSeparateParameters(selector IPodSelector, options *PodSchedulingSeparateOptions) error {
	if selector == nil {
		return fmt.Errorf("parameter selector is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

func (p *jsiiProxy_PodScheduling) validateTolerateParameters(node TaintedNode) error {
	if node == nil {
		return fmt.Errorf("parameter node is required, but nil was provided")
	}

	return nil
}

func validateNewPodSchedulingParameters(instance AbstractPod) error {
	if instance == nil {
		return fmt.Errorf("parameter instance is required, but nil was provided")
	}

	return nil
}

