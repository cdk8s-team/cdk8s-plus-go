//go:build !no_runtime_type_checking

package cdk8splus24

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func validateIngressBackend_FromResourceParameters(resource IResource) error {
	if resource == nil {
		return fmt.Errorf("parameter resource is required, but nil was provided")
	}

	return nil
}

func validateIngressBackend_FromServiceParameters(serv Service, options *ServiceIngressBackendOptions) error {
	if serv == nil {
		return fmt.Errorf("parameter serv is required, but nil was provided")
	}

	if err := _jsii_.ValidateStruct(options, func() string { return "parameter options" }); err != nil {
		return err
	}

	return nil
}

