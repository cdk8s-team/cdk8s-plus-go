package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"
)

// Container environment variables.
type Env interface {
	// The list of sources used to populate the container environment, in addition to the `variables`.
	//
	// Returns a copy. To add a source use `container.env.copyFrom()`.
	Sources() *[]EnvFrom
	// The environment variables for this container.
	//
	// Returns a copy. To add environment variables use `container.env.addVariable()`.
	Variables() *map[string]EnvValue
	// Add a single variable by name and value.
	//
	// The variable value can come from various dynamic sources such a secrets of config maps.
	// Use `EnvValue.fromXXX` to select sources.
	AddVariable(name *string, value EnvValue)
	// Add a collection of variables by copying from another source.
	//
	// Use `Env.fromXXX` functions to select sources.
	CopyFrom(from EnvFrom)
}

// The jsii proxy struct for Env
type jsiiProxy_Env struct {
	_ byte // padding
}

func (j *jsiiProxy_Env) Sources() *[]EnvFrom {
	var returns *[]EnvFrom
	_jsii_.Get(
		j,
		"sources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Env) Variables() *map[string]EnvValue {
	var returns *map[string]EnvValue
	_jsii_.Get(
		j,
		"variables",
		&returns,
	)
	return returns
}


func NewEnv(sources *[]EnvFrom, variables *map[string]EnvValue) Env {
	_init_.Initialize()

	if err := validateNewEnvParameters(sources, variables); err != nil {
		panic(err)
	}
	j := jsiiProxy_Env{}

	_jsii_.Create(
		"cdk8s-plus-30.Env",
		[]interface{}{sources, variables},
		&j,
	)

	return &j
}

func NewEnv_Override(e Env, sources *[]EnvFrom, variables *map[string]EnvValue) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-30.Env",
		[]interface{}{sources, variables},
		e,
	)
}

// Selects a ConfigMap to populate the environment variables with.
//
// The contents of the target ConfigMap's Data field will represent
// the key-value pairs as environment variables.
func Env_FromConfigMap(configMap IConfigMap, prefix *string) EnvFrom {
	_init_.Initialize()

	if err := validateEnv_FromConfigMapParameters(configMap); err != nil {
		panic(err)
	}
	var returns EnvFrom

	_jsii_.StaticInvoke(
		"cdk8s-plus-30.Env",
		"fromConfigMap",
		[]interface{}{configMap, prefix},
		&returns,
	)

	return returns
}

// Selects a Secret to populate the environment variables with.
//
// The contents of the target Secret's Data field will represent
// the key-value pairs as environment variables.
func Env_FromSecret(secr ISecret) EnvFrom {
	_init_.Initialize()

	if err := validateEnv_FromSecretParameters(secr); err != nil {
		panic(err)
	}
	var returns EnvFrom

	_jsii_.StaticInvoke(
		"cdk8s-plus-30.Env",
		"fromSecret",
		[]interface{}{secr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Env) AddVariable(name *string, value EnvValue) {
	if err := e.validateAddVariableParameters(name, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addVariable",
		[]interface{}{name, value},
	)
}

func (e *jsiiProxy_Env) CopyFrom(from EnvFrom) {
	if err := e.validateCopyFromParameters(from); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"copyFrom",
		[]interface{}{from},
	)
}

