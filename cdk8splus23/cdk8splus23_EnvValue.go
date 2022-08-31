// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
)

// Utility class for creating reading env values from various sources.
type EnvValue interface {
	Value() interface{}
	ValueFrom() interface{}
}

// The jsii proxy struct for EnvValue
type jsiiProxy_EnvValue struct {
	_ byte // padding
}

func (j *jsiiProxy_EnvValue) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EnvValue) ValueFrom() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"valueFrom",
		&returns,
	)
	return returns
}


// Create a value by reading a specific key inside a config map.
func EnvValue_FromConfigMap(configMap IConfigMap, key *string, options *EnvValueFromConfigMapOptions) EnvValue {
	_init_.Initialize()

	if err := validateEnvValue_FromConfigMapParameters(configMap, key, options); err != nil {
		panic(err)
	}
	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.EnvValue",
		"fromConfigMap",
		[]interface{}{configMap, key, options},
		&returns,
	)

	return returns
}

// Create a value from a field reference.
func EnvValue_FromFieldRef(fieldPath EnvFieldPaths, options *EnvValueFromFieldRefOptions) EnvValue {
	_init_.Initialize()

	if err := validateEnvValue_FromFieldRefParameters(fieldPath, options); err != nil {
		panic(err)
	}
	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.EnvValue",
		"fromFieldRef",
		[]interface{}{fieldPath, options},
		&returns,
	)

	return returns
}

// Create a value from a key in the current process environment.
func EnvValue_FromProcess(key *string, options *EnvValueFromProcessOptions) EnvValue {
	_init_.Initialize()

	if err := validateEnvValue_FromProcessParameters(key, options); err != nil {
		panic(err)
	}
	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.EnvValue",
		"fromProcess",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

// Create a value from a resource.
func EnvValue_FromResource(resource ResourceFieldPaths, options *EnvValueFromResourceOptions) EnvValue {
	_init_.Initialize()

	if err := validateEnvValue_FromResourceParameters(resource, options); err != nil {
		panic(err)
	}
	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.EnvValue",
		"fromResource",
		[]interface{}{resource, options},
		&returns,
	)

	return returns
}

// Defines an environment value from a secret JSON value.
func EnvValue_FromSecretValue(secretValue *SecretValue, options *EnvValueFromSecretOptions) EnvValue {
	_init_.Initialize()

	if err := validateEnvValue_FromSecretValueParameters(secretValue, options); err != nil {
		panic(err)
	}
	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.EnvValue",
		"fromSecretValue",
		[]interface{}{secretValue, options},
		&returns,
	)

	return returns
}

// Create a value from the given argument.
func EnvValue_FromValue(value *string) EnvValue {
	_init_.Initialize()

	if err := validateEnvValue_FromValueParameters(value); err != nil {
		panic(err)
	}
	var returns EnvValue

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.EnvValue",
		"fromValue",
		[]interface{}{value},
		&returns,
	)

	return returns
}

