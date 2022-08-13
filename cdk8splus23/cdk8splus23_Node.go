// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
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
		"cdk8s-plus-23.Node",
		nil, // no parameters
		&j,
	)

	return &j
}

func NewNode_Override(n Node) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-23.Node",
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
		"cdk8s-plus-23.Node",
		"labeled",
		args,
		&returns,
	)

	return returns
}

// Match a node by its name.
func Node_Named(nodeName *string) NamedNode {
	_init_.Initialize()

	var returns NamedNode

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.Node",
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
		"cdk8s-plus-23.Node",
		"tainted",
		args,
		&returns,
	)

	return returns
}

