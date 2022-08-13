// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"

	"github.com/cdk8s-team/cdk8s-core-go/cdk8s/v2"
)

// Taint queries that can be perfomed against nodes.
type NodeTaintQuery interface {
	Effect() *string
	EvictAfter() cdk8s.Duration
	Key() *string
	Operator() *string
	Value() *string
}

// The jsii proxy struct for NodeTaintQuery
type jsiiProxy_NodeTaintQuery struct {
	_ byte // padding
}

func (j *jsiiProxy_NodeTaintQuery) Effect() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeTaintQuery) EvictAfter() cdk8s.Duration {
	var returns cdk8s.Duration
	_jsii_.Get(
		j,
		"evictAfter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeTaintQuery) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeTaintQuery) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NodeTaintQuery) Value() *string {
	var returns *string
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Matches any taint.
func NodeTaintQuery_Any() NodeTaintQuery {
	_init_.Initialize()

	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NodeTaintQuery",
		"any",
		nil, // no parameters
		&returns,
	)

	return returns
}

// Matches a tain with any value of a specific key.
func NodeTaintQuery_Exists(key *string, options *NodeTaintQueryOptions) NodeTaintQuery {
	_init_.Initialize()

	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NodeTaintQuery",
		"exists",
		[]interface{}{key, options},
		&returns,
	)

	return returns
}

// Matches a taint with a specific key and value.
func NodeTaintQuery_Is(key *string, value *string, options *NodeTaintQueryOptions) NodeTaintQuery {
	_init_.Initialize()

	var returns NodeTaintQuery

	_jsii_.StaticInvoke(
		"cdk8s-plus-22.NodeTaintQuery",
		"is",
		[]interface{}{key, value, options},
		&returns,
	)

	return returns
}

