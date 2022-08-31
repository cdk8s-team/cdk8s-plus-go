// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-22 synthesizes Kubernetes manifests for Kubernetes 1.22.0
package cdk8splus22

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus22/v2/jsii"
)

// A node that is matched by taint selectors.
type TaintedNode interface {
	TaintSelector() *[]NodeTaintQuery
}

// The jsii proxy struct for TaintedNode
type jsiiProxy_TaintedNode struct {
	_ byte // padding
}

func (j *jsiiProxy_TaintedNode) TaintSelector() *[]NodeTaintQuery {
	var returns *[]NodeTaintQuery
	_jsii_.Get(
		j,
		"taintSelector",
		&returns,
	)
	return returns
}


func NewTaintedNode(taintSelector *[]NodeTaintQuery) TaintedNode {
	_init_.Initialize()

	if err := validateNewTaintedNodeParameters(taintSelector); err != nil {
		panic(err)
	}
	j := jsiiProxy_TaintedNode{}

	_jsii_.Create(
		"cdk8s-plus-22.TaintedNode",
		[]interface{}{taintSelector},
		&j,
	)

	return &j
}

func NewTaintedNode_Override(t TaintedNode, taintSelector *[]NodeTaintQuery) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-22.TaintedNode",
		[]interface{}{taintSelector},
		t,
	)
}

