package cdk8splus27

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdk8s-team/cdk8s-plus-go/cdk8splus27/v2/jsii"
)

// A metric condition that HorizontalPodAutoscaler's scale on.
type Metric interface {
	Type() *string
}

// The jsii proxy struct for Metric
type jsiiProxy_Metric struct {
	_ byte // padding
}

func (j *jsiiProxy_Metric) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}


// Metric that tracks the CPU of a container.
//
// This metric
// will be tracked across all pods of the current scale target.
func Metric_ContainerCpu(options *MetricContainerResourceOptions) Metric {
	_init_.Initialize()

	if err := validateMetric_ContainerCpuParameters(options); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"containerCpu",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Metric that tracks the local ephemeral storage of a container.
//
// This metric
// will be tracked across all pods of the current scale target.
func Metric_ContainerEphemeralStorage(options *MetricContainerResourceOptions) Metric {
	_init_.Initialize()

	if err := validateMetric_ContainerEphemeralStorageParameters(options); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"containerEphemeralStorage",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Metric that tracks the Memory of a container.
//
// This metric
// will be tracked across all pods of the current scale target.
func Metric_ContainerMemory(options *MetricContainerResourceOptions) Metric {
	_init_.Initialize()

	if err := validateMetric_ContainerMemoryParameters(options); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"containerMemory",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Metric that tracks the volume size of a container.
//
// This metric
// will be tracked across all pods of the current scale target.
func Metric_ContainerStorage(options *MetricContainerResourceOptions) Metric {
	_init_.Initialize()

	if err := validateMetric_ContainerStorageParameters(options); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"containerStorage",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// A global metric that is not associated with any Kubernetes object.
//
// Allows for autoscaling based on information coming from components running outside of
// the cluster.
//
// Use case:
// * Scale up when the length of an SQS queue is greater than 10 messages.
// * Scale down when an outside load balancer's queries are less than 10000 per second.
func Metric_External(options *MetricOptions) Metric {
	_init_.Initialize()

	if err := validateMetric_ExternalParameters(options); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"external",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Metric that describes a metric of a kubernetes object.
//
// Use case:
// * Scale on a Kubernetes Ingress's hits-per-second metric.
func Metric_Object(options *MetricObjectOptions) Metric {
	_init_.Initialize()

	if err := validateMetric_ObjectParameters(options); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"object",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// A pod metric that will be averaged across all pods of the current scale target.
//
// Use case:
// * Average CPU utilization across all pods
// * Transactions processed per second across all pods.
func Metric_Pods(options *MetricOptions) Metric {
	_init_.Initialize()

	if err := validateMetric_PodsParameters(options); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"pods",
		[]interface{}{options},
		&returns,
	)

	return returns
}

// Tracks the available CPU of the pods in a target.
//
// Note: Since the resource usages of all the containers are summed up the total
// pod utilization may not accurately represent the individual container resource
// usage. This could lead to situations where a single container might be running
// with high usage and the HPA will not scale out because the overall pod usage
// is still within acceptable limits.
//
// Use case:
// * Scale up when CPU is above 40%.
func Metric_ResourceCpu(target MetricTarget) Metric {
	_init_.Initialize()

	if err := validateMetric_ResourceCpuParameters(target); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"resourceCpu",
		[]interface{}{target},
		&returns,
	)

	return returns
}

// Tracks the available Ephemeral Storage of the pods in a target.
//
// Note: Since the resource usages of all the containers are summed up the total
// pod utilization may not accurately represent the individual container resource
// usage. This could lead to situations where a single container might be running
// with high usage and the HPA will not scale out because the overall pod usage
// is still within acceptable limits.
func Metric_ResourceEphemeralStorage(target MetricTarget) Metric {
	_init_.Initialize()

	if err := validateMetric_ResourceEphemeralStorageParameters(target); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"resourceEphemeralStorage",
		[]interface{}{target},
		&returns,
	)

	return returns
}

// Tracks the available Memory of the pods in a target.
//
// Note: Since the resource usages of all the containers are summed up the total
// pod utilization may not accurately represent the individual container resource
// usage. This could lead to situations where a single container might be running
// with high usage and the HPA will not scale out because the overall pod usage
// is still within acceptable limits.
//
// Use case:
// * Scale up when Memory is above 512MB.
func Metric_ResourceMemory(target MetricTarget) Metric {
	_init_.Initialize()

	if err := validateMetric_ResourceMemoryParameters(target); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"resourceMemory",
		[]interface{}{target},
		&returns,
	)

	return returns
}

// Tracks the available Storage of the pods in a target.
//
// Note: Since the resource usages of all the containers are summed up the total
// pod utilization may not accurately represent the individual container resource
// usage. This could lead to situations where a single container might be running
// with high usage and the HPA will not scale out because the overall pod usage
// is still within acceptable limits.
func Metric_ResourceStorage(target MetricTarget) Metric {
	_init_.Initialize()

	if err := validateMetric_ResourceStorageParameters(target); err != nil {
		panic(err)
	}
	var returns Metric

	_jsii_.StaticInvoke(
		"cdk8s-plus-27.Metric",
		"resourceStorage",
		[]interface{}{target},
		&returns,
	)

	return returns
}

