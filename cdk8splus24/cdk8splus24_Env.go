// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"
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

	j := jsiiProxy_Env{}

	_jsii_.Create(
		"cdk8s-plus-24.Env",
		[]interface{}{sources, variables},
		&j,
	)

	return &j
}

func NewEnv_Override(e Env, sources *[]EnvFrom, variables *map[string]EnvValue) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Env",
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

	var returns EnvFrom

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Env",
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

	var returns EnvFrom

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Env",
		"fromSecret",
		[]interface{}{secr},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Env) AddVariable(name *string, value EnvValue) {
	_jsii_.InvokeVoid(
		e,
		"addVariable",
		[]interface{}{name, value},
	)
}

func (e *jsiiProxy_Env) CopyFrom(from EnvFrom) {
	_jsii_.InvokeVoid(
		e,
		"copyFrom",
		[]interface{}{from},
	)
}

