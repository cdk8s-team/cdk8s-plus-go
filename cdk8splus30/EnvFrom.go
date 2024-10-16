package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"
)

// A collection of env variables defined in other resources.
type EnvFrom interface {
}

// The jsii proxy struct for EnvFrom
type jsiiProxy_EnvFrom struct {
	_ byte // padding
}

func NewEnvFrom(configMap IConfigMap, prefix *string, sec ISecret) EnvFrom {
	_init_.Initialize()

	j := jsiiProxy_EnvFrom{}

	_jsii_.Create(
		"cdk8s-plus-30.EnvFrom",
		[]interface{}{configMap, prefix, sec},
		&j,
	)

	return &j
}

func NewEnvFrom_Override(e EnvFrom, configMap IConfigMap, prefix *string, sec ISecret) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-30.EnvFrom",
		[]interface{}{configMap, prefix, sec},
		e,
	)
}

