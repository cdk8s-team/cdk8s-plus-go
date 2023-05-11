package cdk8splus24

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus24/v2/jsii"
)

// A metric condition that will trigger scaling behavior when satisfied.
//
// Example:
//   MetricTarget.averageUtilization(70); // 70% average utilization
//
type MetricTarget interface {
}

// The jsii proxy struct for MetricTarget
type jsiiProxy_MetricTarget struct {
	_ byte // padding
}

// Target a percentage value across all relevant pods.
func MetricTarget_AverageUtilization(averageUtilization *float64) MetricTarget {
	_init_.Initialize()

	if err := validateMetricTarget_AverageUtilizationParameters(averageUtilization); err != nil {
		panic(err)
	}
	var returns MetricTarget

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.MetricTarget",
		"averageUtilization",
		[]interface{}{averageUtilization},
		&returns,
	)

	return returns
}

// Target the average value across all relevant pods.
func MetricTarget_AverageValue(averageValue *float64) MetricTarget {
	_init_.Initialize()

	if err := validateMetricTarget_AverageValueParameters(averageValue); err != nil {
		panic(err)
	}
	var returns MetricTarget

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.MetricTarget",
		"averageValue",
		[]interface{}{averageValue},
		&returns,
	)

	return returns
}

// Target a specific target value.
func MetricTarget_Value(value *float64) MetricTarget {
	_init_.Initialize()

	if err := validateMetricTarget_ValueParameters(value); err != nil {
		panic(err)
	}
	var returns MetricTarget

	_jsii_.StaticInvoke(
		"cdk8s-plus-24.MetricTarget",
		"value",
		[]interface{}{value},
		&returns,
	)

	return returns
}

