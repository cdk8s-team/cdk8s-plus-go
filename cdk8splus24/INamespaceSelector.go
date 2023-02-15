// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/internal"
)

// Represents an object that can select namespaces.
type INamespaceSelector interface {
	constructs.IConstruct
	// Return the configuration of this selector.
	ToNamespaceSelectorConfig() *NamespaceSelectorConfig
}

// The jsii proxy for INamespaceSelector
type jsiiProxy_INamespaceSelector struct {
	internal.Type__constructsIConstruct
}

func (i *jsiiProxy_INamespaceSelector) ToNamespaceSelectorConfig() *NamespaceSelectorConfig {
	var returns *NamespaceSelectorConfig

	_jsii_.Invoke(
		i,
		"toNamespaceSelectorConfig",
		nil, // no parameters
		&returns,
	)

	return returns
}

