// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
)

// Taint queries that can be perfomed against nodes.
type NodeTaintQuery interface {
}

// The jsii proxy struct for NodeTaintQuery
type jsiiProxy_NodeTaintQuery struct {
	_ byte // padding
}

// Matches any taint.
func NodeTaintQuery_Any() NodeTaintQuery {
	_init_.Initialize()

	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.NodeTaintQuery",
		"any",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches a tain with any value of a specific key.
func NodeTaintQuery_Exists(key *string, options *NodeTaintQueryOptions) NodeTaintQuery {
	_init_.Initialize()

	if err := validateNodeTaintQuery_ExistsParameters(key, options); err != nil {
		panic(err)
	}
	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.NodeTaintQuery",
		"exists",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

// Matches a taint with a specific key and value.
func NodeTaintQuery_Is(key *string, value *string, options *NodeTaintQueryOptions) NodeTaintQuery {
	_init_.Initialize()

	if err := validateNodeTaintQuery_IsParameters(key, value, options); err != nil {
		panic(err)
	}
	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.NodeTaintQuery",
		"is",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

