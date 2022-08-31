//go:build !no_runtime_type_checking
// +build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
)

func (i *jsiiProxy_Ingress) validateAddDefaultBackendParameters(backend IngressBackend) error {
	if backend == nil {
		return fmt.Errorf("parameter backend is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Ingress) validateAddHostDefaultBackendParameters(host *string, backend IngressBackend) error {
	if host == nil {
		return fmt.Errorf("parameter host is required, but nil was provided")
	}

	if backend == nil {
		return fmt.Errorf("parameter backend is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Ingress) validateAddHostRuleParameters(host *string, path *string, backend IngressBackend) error {
	if host == nil {
		return fmt.Errorf("parameter host is required, but nil was provided")
	}

	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if backend == nil {
		return fmt.Errorf("parameter backend is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Ingress) validateAddRuleParameters(path *string, backend IngressBackend) error {
	if path == nil {
		return fmt.Errorf("parameter path is required, but nil was provided")
	}

	if backend == nil {
		return fmt.Errorf("parameter backend is required, but nil was provided")
	}

	return nil
}

func (i *jsiiProxy_Ingress) validateAddRulesParameters(rules *[]*IngressRule) error {
	for idx_6c621d, v := range *rules {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter rules[%#v]", idx_6c621d) }); err != nil {
			return err
		}
	}

	return nil
}

func (i *jsiiProxy_Ingress) validateAddTlsParameters(tls *[]*IngressTls) error {
	if tls == nil {
		return fmt.Errorf("parameter tls is required, but nil was provided")
	}
	for idx_b7e651, v := range *tls {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter tls[%#v]", idx_b7e651) }); err != nil {
			return err
		}
	}

	return nil
}

func validateIngress_IsConstructParameters(x interface{}) error {
	if x == nil {
		return fmt.Errorf("parameter x is required, but nil was provided")
	}

	return nil
}

func validateNewIngressParameters(scope constructs.Construct, id *string, props *IngressProps) error {
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

