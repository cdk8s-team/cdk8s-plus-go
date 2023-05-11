package cdk8splus25

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

