// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-26 synthesizes Kubernetes manifests for Kubernetes 1.26.0
package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/jsii"
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
		"cdk8s-plus-26.EnvFrom",
		[]interface{}{configMap, prefix, sec},
		&j,
	)

	return &j
}

func NewEnvFrom_Override(e EnvFrom, configMap IConfigMap, prefix *string, sec ISecret) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-26.EnvFrom",
		[]interface{}{configMap, prefix, sec},
		e,
	)
}

