package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"
)

// Represents a node in the cluster.
type Node interface {
}

// The jsii proxy struct for Node
type jsiiProxy_Node struct {
	_ byte // padding
}

func NewNode() Node {
	_init_.Initialize()

	j := jsiiProxy_Node{}

	_jsii_.Create(
		"cdk8s-plus-24.Node",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNode_Override(n Node) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-24.Node",
		nil, // no parameters
		n,
	)
}

// Match a node by its labels.
func Node_Labeled(labelSelector ...NodeLabelQuery) LabeledNode {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range labelSelector {
		args = append(args, a)
	}

	var returns LabeledNode

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Node",
		"labeled",
		args,
		&returns,
	)

	return returns
}

// Match a node by its name.
func Node_Named(nodeName *string) NamedNode {
	_init_.Initialize()

	if err := validateNode_NamedParameters(nodeName); err != nil {
		panic(err)
	}
	var returns NamedNode

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Node",
		"named",
		[]interface{}{nodeName},
		&returns,
	)

	return returns
}

// Match a node by its taints.
func Node_Tainted(taintSelector ...NodeTaintQuery) TaintedNode {
	_init_.Initialize()

	args := []interface{}{}
	for _, a := range taintSelector {
		args = append(args, a)
	}

	var returns TaintedNode

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.Node",
		"tainted",
		args,
		&returns,
	)

	return returns
}

