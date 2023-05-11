//go:build !no_runtime_type_checking

package cdk8splus25

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (c *jsiiProxy_CronJob) validateAddContainerParameters(cont *ContainerProps) error {
	if cont == nil {
		return fmt.Errorf("parameter cont is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cont, func() string { return "parameter cont" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CronJob) validateAddHostAliasParameters(hostAlias *HostAlias) error {
	if hostAlias == nil {
		return fmt.Errorf("parameter hostAlias is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(hostAlias, func() string { return "parameter hostAlias" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CronJob) validateAddInitContainerParameters(cont *ContainerProps) error {
	if cont == nil {
		return fmt.Errorf("parameter cont is required, but nil was provided")
	}
	if err := _jsii_.ValidateStruct(cont, func() string { return "parameter cont" }); err != nil {
		return err
	}

	return nil
}

func (c *jsiiProxy_CronJob) validateAddVolumeParameters(vol Volume) error {
	if vol == nil {
		return fmt.Errorf("parameter vol is required, but nil was provided")
	}

	return nil
}

func (c *jsiiProxy_CronJob) validateAttachContainerParameters(cont Container) error {
	if cont == nil {
		return fmt.Errorf("parameter cont is required, but nil was provided")
	}

	return nil
}

func validateCronJob_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewCronJobParameters(scope constructs.Construct, id *string, props *CronJobProps) error {
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

