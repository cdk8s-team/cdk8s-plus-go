// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
)

// Represents a query that can be performed against resources with labels.
type LabelExpression interface {
	Key() *string
	Operator() *string
	Values() *[]*string
}

// The jsii proxy struct for LabelExpression
type jsiiProxy_LabelExpression struct {
	_ byte // padding
}

func (j *jsiiProxy_LabelExpression) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabelExpression) Operator() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabelExpression) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}


// Requires label `key` to not exist.
func LabelExpression_DoesNotExist(key *string) LabelExpression {
	_init_.Initialize()

	var returns LabelExpression

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.LabelExpression",
		"doesNotExist",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Requires label `key` to exist.
func LabelExpression_Exists(key *string) LabelExpression {
	_init_.Initialize()

	var returns LabelExpression

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.LabelExpression",
		"exists",
		[]interface{}{key},
		&returns,
	)

	return returns
}

// Requires value of label `key` to be one of `values`.
func LabelExpression_In(key *string, values *[]*string) LabelExpression {
	_init_.Initialize()

	var returns LabelExpression

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.LabelExpression",
		"in",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

// Requires value of label `key` to be none of `values`.
func LabelExpression_NotIn(key *string, values *[]*string) LabelExpression {
	_init_.Initialize()

	var returns LabelExpression

	_jsii_.StaticInvoke(
		"cdk8s-plus-23.LabelExpression",
		"notIn",
		[]interface{}{key, values},
		&returns,
	)

	return returns
}

