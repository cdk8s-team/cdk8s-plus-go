package cdk8splus26

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus26/v2/jsii"
)

// Union like class repsenting either a ration in percents or an absolute number.
type PercentOrAbsolute interface {
	Value() interface{}
	IsZero() *bool
}

// The jsii proxy struct for PercentOrAbsolute
type jsiiProxy_PercentOrAbsolute struct {
	_ byte // padding
}

func (j *jsiiProxy_PercentOrAbsolute) Value() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"value",
		&returns,
	)
	return returns
}


// Absolute number.
func PercentOrAbsolute_Absolute(num *float64) PercentOrAbsolute {
	_init_.Initialize()

	if err := validatePercentOrAbsolute_AbsoluteParameters(num); err != nil {
		panic(err)
	}
	var returns PercentOrAbsolute

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.PercentOrAbsolute",
		"absolute",
		[]interface{}{num},
		&returns,
	)

	return returns
}

// Percent ratio.
func PercentOrAbsolute_Percent(percent *float64) PercentOrAbsolute {
	_init_.Initialize()

	if err := validatePercentOrAbsolute_PercentParameters(percent); err != nil {
		panic(err)
	}
	var returns PercentOrAbsolute

	_jsii_.StaticInvoke(
		"cdk8s-plus-26.PercentOrAbsolute",
		"percent",
		[]interface{}{percent},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PercentOrAbsolute) IsZero() *bool {
	var returns *bool

	_jsii_.Invoke(
		p,
		"isZero",
		nil, // no parameters
		&returns,
	)

	return returns
}

