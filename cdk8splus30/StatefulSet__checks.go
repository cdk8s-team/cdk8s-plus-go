//go:build !no_runtime_type_checking

package cdk8splus30

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (s *jsiiProxy_StatefulSet) validateAddContainerParameters(cont *ContainerProps) error {
	if cont == nil {
		return fmt.Errorf("parameter cont is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cont, func() string { return "parameter cont" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSet) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	if hostAlias == nil {
		return fmt.Errorf("parameter hostAlias is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(hostAlias, func() string { return "parameter hostAlias" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSet) validateAddInitContainerParameters(cont *ContainerProps) error {
	if cont == nil {
		return fmt.Errorf("parameter cont is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cont, func() string { return "parameter cont" }); err != nil {
		return err
	}

	return nil
}

func (s *jsiiProxy_StatefulSet) validateAddVolumeParameters(vol Volume) error {
	if vol == nil {
		return fmt.Errorf("parameter vol is required, but nil was provided")
	}

	return nil
}

func (s *jsiiProxy_StatefulSet) validateAttachContainerParameters(cont Container) error {
	if cont == nil {
		return fmt.Errorf("parameter cont is required, but nil was provided")
	}

	return nil
}

func validateStatefulSet_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func (j *jsiiProxy_StatefulSet) validateSetHasAutoscalerParameters(val *bool) error {
	if val == nil {
		return fmt.Errorf("parameter val is required, but nil was provided")
	}

	return nil
}

func validateNewStatefulSetParameters(scope constructs.Construct, id *string, props *StatefulSetProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if props == nil {
		return fmt.Errorf("parameter props is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

