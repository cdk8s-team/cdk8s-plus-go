//go:build no_runtime_type_checking

package cdk8splus30

// Building without runtime type checking enabled, so all the below just return nil

func validateHandler_FromCommandParameters(command *[]*string) error {
	return nil
}

func validateHandler_FromHttpGetParameters(path *string, options *HandlerFromHttpGetOptions) error {
	return nil
}

func validateHandler_FromTcpSocketParameters(options *HandlerFromTcpSocketOptions) error {
	return nil
}

