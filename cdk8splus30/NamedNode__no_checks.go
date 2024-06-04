//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func validateNewNamedNodeParameters(name *string) error {
	return nil
}

