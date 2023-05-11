//go:build no_runtime_type_checking

package cdk8splus24

// Building without runtime type checking enabled, so all the below just return nil

func validateDeploymentStrategy_RollingUpdateParameters(options *DeploymentStrategyRollingUpdateOptions) error {
	return nil
}

