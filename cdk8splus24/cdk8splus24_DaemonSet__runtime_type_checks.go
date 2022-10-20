//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (d *jsiiProxy_DaemonSet) validateAddContainerParameters(cont *ContainerProps) error {
	if cont == nil {
		return fmt.Errorf("parameter cont is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cont, func() string { return "parameter cont" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DaemonSet) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	if hostAlias == nil {
		return fmt.Errorf("parameter hostAlias is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(hostAlias, func() string { return "parameter hostAlias" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DaemonSet) validateAddInitContainerParameters(cont *ContainerProps) error {
	if cont == nil {
		return fmt.Errorf("parameter cont is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cont, func() string { return "parameter cont" }); err != nil {
		return err
	}

	return nil
}

func (d *jsiiProxy_DaemonSet) validateAddVolumeParameters(vol Volume) error {
	if vol == nil {
		return fmt.Errorf("parameter vol is required, but nil was provided")
	}

	return nil
}

func validateDaemonSet_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewDaemonSetParameters(scope constructs.Construct, id *string, props *DaemonSetProps) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

