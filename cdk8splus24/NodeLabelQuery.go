// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-24 synthesizes Kubernetes manifests for Kubernetes 1.24.0
package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"
)

// Represents a query that can be performed against nodes with labels.
type NodeLabelQuery interface {
}

// The jsii proxy struct for NodeLabelQuery
type jsiiProxy_NodeLabelQuery struct {
	_ byte // padding
}

// Requires label `key` to not exist.
func NodeLabelQuery_DoesNotExist(key *string) NodeLabelQuery {
	_init_.Initialize()

	if err := validateNodeLabelQuery_DoesNotExistParameters(key); err != nil {
		panic(err)
	}
	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"doesNotExist",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Requires label `key` to exist.
func NodeLabelQuery_Exists(key *string) NodeLabelQuery {
	_init_.Initialize()

	if err := validateNodeLabelQuery_ExistsParameters(key); err != nil {
		panic(err)
	}
	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"exists",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Requires value of label `key` to greater than all elements in `values`.
func NodeLabelQuery_Gt(key *string, values *[]*string) NodeLabelQuery {
	_init_.Initialize()

	if err := validateNodeLabelQuery_GtParameters(key, values); err != nil {
		panic(err)
	}
	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"gt",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Requires value of label `key` to be one of `values`.
func NodeLabelQuery_In(key *string, values *[]*string) NodeLabelQuery {
	_init_.Initialize()

	if err := validateNodeLabelQuery_InParameters(key, values); err != nil {
		panic(err)
	}
	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"in",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Requires value of label `key` to equal `value`.
func NodeLabelQuery_Is(key *string, value *string) NodeLabelQuery {
	_init_.Initialize()

	if err := validateNodeLabelQuery_IsParameters(key, value); err != nil {
		panic(err)
	}
	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"is",
		[]interface{}{key, value},
		&returns,
	)

	return returns
}

// Requires value of label `key` to less than all elements in `values`.
func NodeLabelQuery_Lt(key *string, values *[]*string) NodeLabelQuery {
	_init_.Initialize()

	if err := validateNodeLabelQuery_LtParameters(key, values); err != nil {
		panic(err)
	}
	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"lt",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Requires value of label `key` to be none of `values`.
func NodeLabelQuery_NotIn(key *string, values *[]*string) NodeLabelQuery {
	_init_.Initialize()

	if err := validateNodeLabelQuery_NotInParameters(key, values); err != nil {
		panic(err)
	}
	var returns NodeLabelQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.NodeLabelQuery",
		"notIn",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

