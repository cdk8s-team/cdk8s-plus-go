package cdk8splus31

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus31/v2/jsii"
)

// Represents the amount of CPU.
//
// The amount can be passed as millis or units.
type Cpu interface {
	Amount() *string
	SetAmount(val *string)
}

// The jsii proxy struct for Cpu
type jsiiProxy_Cpu struct {
	_ byte // padding
}

func (j *jsiiProxy_Cpu) Amount() *string {
	var returns *string
	_jsii_.Get(
		j,
		"amount",
		&returns,
	)
	return returns
}


func (j *jsiiProxy_Cpu)SetAmount(val *string) {
	if err := j.validateSetAmountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"amount",
		val,
	)
}

func Cpu_Millis(amount *float64) Cpu {
	_init_.Initialize()

	if err := validateCpu_MillisParameters(amount); err != nil {
		panic(err)
	}
	var returns Cpu

	_jsii_.StaticInvoke(
		"cdk8s-plus-31.Cpu",
		"millis",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

func Cpu_Units(amount *float64) Cpu {
	_init_.Initialize()

	if err := validateCpu_UnitsParameters(amount); err != nil {
		panic(err)
	}
	var returns Cpu

	_jsii_.StaticInvoke(
		"cdk8s-plus-31.Cpu",
		"units",
		[]interface{}{amount},
		&returns,
	)

	return returns
}

