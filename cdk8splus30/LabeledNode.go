package cdk8splus30

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus30/v2/jsii"
)

// A node that is matched by label selectors.
type LabeledNode interface {
	LabelSelector() *[]NodeLabelQuery
}

// The jsii proxy struct for LabeledNode
type jsiiProxy_LabeledNode struct {
	_ byte // padding
}

func (j *jsiiProxy_LabeledNode) LabelSelector() *[]NodeLabelQuery {
	var returns *[]NodeLabelQuery
	_jsii_.Get(
		j,
		"labelSelector",
		&returns,
	)
	return returns
}


func NewLabeledNode(labelSelector *[]NodeLabelQuery) LabeledNode {
	_init_.Initialize()

	if err := validateNewLabeledNodeParameters(labelSelector); err != nil {
		panic(err)
	}
	j := jsiiProxy_LabeledNode{}

	_jsii_.Create(
		"cdk8s-plus-30.LabeledNode",
		[]interface{}{labelSelector},
		&j,
	)

	return &j
}

func NewLabeledNode_Override(l LabeledNode, labelSelector *[]NodeLabelQuery) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-30.LabeledNode",
		[]interface{}{labelSelector},
		l,
	)
}

