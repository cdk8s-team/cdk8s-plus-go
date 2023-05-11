//go:build no_runtime_type_checking

package cdk8splus25

// Building without runtime type checking enabled, so all the below just return nil

func validateApiResource_CustomParameters(options *ApiResourceOptions) error {
	return nil
}

