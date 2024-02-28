//go:build no_runtime_type_checking

package cdk8splus28

// Building without runtime type checking enabled, so all the below just return nil

func validateNewContainerSecurityContextParameters(props *ContainerSecurityContextProps) error {
	return nil
}

