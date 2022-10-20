//go:build !no_runtime_type_checking

// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	"fmt"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func (p *jsiiProxy_PodDns) validateAddOptionParameters(options *[]*DnsOption) error {
	for idx_a793ab, v := range *options {
		if err := _jsii_.ValidateStruct(v, func() string { return fmt.Sprintf("parameter options[%#v]", idx_a793ab) }); err != nil {
			return err
		}
	}

	return nil
}

func validateNewPodDnsParameters(props *PodDnsProps) error {
	if err := _jsii_.ValidateStruct(props, func() string { return "parameter props" }); err != nil {
		return err
	}

	return nil
}

