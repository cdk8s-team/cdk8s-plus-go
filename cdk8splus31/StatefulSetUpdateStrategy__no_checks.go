//go:build no_runtime_type_checking

package cdk8splus31

// Building without runtime type checking enabled, so all the below just return nil

func validateStatefulSetUpdateStrategy_RollingUpdateParameters(options *StatefulSetUpdateStrategyRollingUpdateOptions) error {
	return nil
}

