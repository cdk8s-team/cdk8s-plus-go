//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func validateMetric_ContainerCpuParameters(options *MetricContainerResourceOptions) error {
	return nil
}

func validateMetric_ContainerEphemeralStorageParameters(options *MetricContainerResourceOptions) error {
	return nil
}

func validateMetric_ContainerMemoryParameters(options *MetricContainerResourceOptions) error {
	return nil
}

func validateMetric_ContainerStorageParameters(options *MetricContainerResourceOptions) error {
	return nil
}

func validateMetric_ExternalParameters(options *MetricOptions) error {
	return nil
}

func validateMetric_ObjectParameters(options *MetricObjectOptions) error {
	return nil
}

func validateMetric_PodsParameters(options *MetricOptions) error {
	return nil
}

func validateMetric_ResourceCpuParameters(target MetricTarget) error {
	return nil
}

func validateMetric_ResourceEphemeralStorageParameters(target MetricTarget) error {
	return nil
}

func validateMetric_ResourceMemoryParameters(target MetricTarget) error {
	return nil
}

func validateMetric_ResourceStorageParameters(target MetricTarget) error {
	return nil
}

