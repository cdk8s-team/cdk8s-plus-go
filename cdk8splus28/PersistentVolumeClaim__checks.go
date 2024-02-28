//go:build !no_runtime_type_checking

package cdk8splus28

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (p *jsiiProxy_PersistentVolumeClaim) validateBindParameters(vol IPersistentVolume) error {
	if vol == nil {
		return fmt.Errorf("parameter vol is required, but nil was provided")
	}

	return nil
}

func validatePersistentVolumeClaim_FromClaimNameParameters(scope constructs.Construct, id *string, claimName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if claimName == nil {
		return fmt.Errorf("parameter claimName is required, but nil was provided")
	}

	return nil
}

func validatePersistentVolumeClaim_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewPersistentVolumeClaimParameters(scope constructs.Construct, id *string, props *PersistentVolumeClaimProps) error {
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

