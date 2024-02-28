package cdk8splus28

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus28/v2/jsii"
)

// A node that is matched by its name.
type NamedNode interface {
	Name() *string
}

// The jsii proxy struct for NamedNode
type jsiiProxy_NamedNode struct {
	_ byte // padding
}

func (j *jsiiProxy_NamedNode) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}


func NewNamedNode(name *string) NamedNode {
	_init_.Initialize()

	if err := validateNewNamedNodeParameters(name); err != nil {
		panic(err)
	}
	j := jsiiProxy_NamedNode{}

	_jsii_.Create(
		"cdk8s-plus-28.NamedNode",
		[]interface{}{name},
		&j,
	)

	return &j
}

func NewNamedNode_Override(n NamedNode, name *string) {
	_init_.Initialize()

	_jsii_.Create(
		"cdk8s-plus-28.NamedNode",
		[]interface{}{name},
		n,
	)
}

