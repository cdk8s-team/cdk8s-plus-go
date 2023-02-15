// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

type ISecret interface {
	IResource
	// Returns EnvValue object from a secret's key.
	EnvValue(key *string, options *EnvValueFromSecretOptions) EnvValue
}

// The jsii proxy for ISecret
type jsiiProxy_ISecret struct {
	jsiiProxy_IResource
}

func (i *jsiiProxy_ISecret) EnvValue(key *string, options *EnvValueFromSecretOptions) EnvValue {
	if err := i.validateEnvValueParameters(key, options); err != nil {
		panic(err)
	}
	var returns EnvValue

	_jsii_.Invoke(
		i,
		"envValue",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

