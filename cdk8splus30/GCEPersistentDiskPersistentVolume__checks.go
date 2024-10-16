//go:build !no_runtime_type_checking

package cdk8splus30

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (g *jsiiProxy_GCEPersistentDiskPersistentVolume) validateBindParameters(claim IPersistentVolumeClaim) error {
	if claim == nil {
		return fmt.Errorf("parameter claim is required, but nil was provided")
	}

	return nil
}

func validateGCEPersistentDiskPersistentVolume_FromPersistentVolumeNameParameters(scope constructs.Construct, id *string, volumeName *string) error {
	if scope == nil {
		return fmt.Errorf("parameter scope is required, but nil was provided")
	}

	if id == nil {
		return fmt.Errorf("parameter id is required, but nil was provided")
	}

	if volumeName == nil {
		return fmt.Errorf("parameter volumeName is required, but nil was provided")
	}

	return nil
}

func validateGCEPersistentDiskPersistentVolume_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewGCEPersistentDiskPersistentVolumeParameters(scope constructs.Construct, id *string, props *GCEPersistentDiskPersistentVolumeProps) error {
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

