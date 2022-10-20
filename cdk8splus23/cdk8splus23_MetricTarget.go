// cdk8s+ is a software development framework that provides high level abstractions for authoring Kubernetes applications. cdk8s-plus-23 synthesizes Kubernetes manifests for Kubernetes 1.23.0
package cdk8splus23

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus23/v2/jsii"
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
		"cdk8s-plus-23.MetricTarget",
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
		"cdk8s-plus-23.MetricTarget",
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
		"cdk8s-plus-23.MetricTarget",
		"value",
		[]interface{}{value},
		&returns,
	)

	return returns
}

